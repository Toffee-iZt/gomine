package packet

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"

	"github.com/Toffee-iZt/gomine/proto/types"
)

// Packet define a net data package
type Packet struct {
	ID   int32
	Data []byte
}

// Make generates Packet with the ID and Fields
func Make(id int32, fields ...types.Encoder) (pk Packet) {
	var pb Builder
	for _, v := range fields {
		pb.WriteField(v)
	}
	return pb.Packet(id)
}

// Scan decode the packet and fill data into fields
func (p Packet) Scan(fields ...types.Decoder) error {
	r := bytes.NewReader(p.Data)
	for _, v := range fields {
		_, err := v.ReadFrom(r)
		if err != nil {
			return err
		}
	}
	return nil
}

// Pack ...
func (p *Packet) Pack(w io.Writer, threshold int) error {
	var content bytes.Buffer
	if _, err := types.VarInt(p.ID).WriteTo(&content); err != nil {
		panic(err)
	}
	if _, err := content.Write(p.Data); err != nil {
		panic(err)
	}
	if threshold > 0 {
		rawLen := content.Len()
		uncompressedLen := types.VarInt(rawLen)
		if rawLen > threshold {
			compress(&content)
		} else {
			uncompressedLen = 0
		}

		uncompressedLenLen, _ := uncompressedLen.WriteTo(io.Discard)
		if _, err := types.VarInt(uncompressedLenLen + int64(rawLen)).WriteTo(w); err != nil {
			return err
		}

		if _, err := uncompressedLen.WriteTo(w); err != nil {
			return err
		}
		if _, err := content.WriteTo(w); err != nil {
			return err
		}
	} else {
		if _, err := types.VarInt(content.Len()).WriteTo(w); err != nil {
			return err
		}
		if _, err := content.WriteTo(w); err != nil {
			return err
		}
	}

	return nil
}

// UnPack in-place decompression a packet
func (p *Packet) UnPack(r io.Reader, threshold int) error {
	var length types.VarInt
	if _, err := length.ReadFrom(r); err != nil {
		return err
	}
	if length < 1 {
		return fmt.Errorf("packet length too short")
	}
	buf := make([]byte, length)
	if _, err := io.ReadFull(r, buf); err != nil {
		return fmt.Errorf("read content of packet fail: %w", err)
	}
	buffer := bytes.NewBuffer(buf)

	if threshold > 0 {
		if err := unCompress(buffer); err != nil {
			return err
		}
	}

	var packetID types.VarInt
	if _, err := packetID.ReadFrom(buffer); err != nil {
		return fmt.Errorf("read packet id fail: %v", err)
	}
	p.ID = int32(packetID)
	p.Data = buffer.Bytes()
	return nil
}

func unCompress(data *bytes.Buffer) error {
	reader := bytes.NewReader(data.Bytes())

	var sizeUncompressed types.VarInt
	if _, err := sizeUncompressed.ReadFrom(reader); err != nil {
		return err
	}

	var uncompressedData []byte
	if sizeUncompressed != 0 { // != 0 means compressed, let's decompress
		uncompressedData = make([]byte, sizeUncompressed)
		r, err := zlib.NewReader(reader)
		if err != nil {
			return fmt.Errorf("decompress fail: %v", err)
		}
		defer r.Close()
		_, err = io.ReadFull(r, uncompressedData)
		if err != nil {
			return fmt.Errorf("decompress fail: %v", err)
		}
	} else {
		uncompressedData = data.Bytes()[1:]
	}
	*data = *bytes.NewBuffer(uncompressedData)
	return nil
}

func compress(data *bytes.Buffer) {
	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	if _, err := data.WriteTo(w); err != nil {
		panic(err)
	}
	if err := w.Close(); err != nil {
		panic(err)
	}
	*data = b
	return
}
