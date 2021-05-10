package packet

import (
	"bytes"

	"github.com/Toffee-iZt/gomine/proto/types"
)

// Builder ...
type Builder struct {
	buf bytes.Buffer
}

// WriteField ...
func (p *Builder) WriteField(fields ...types.Encoder) {
	for _, f := range fields {
		_, err := f.WriteTo(&p.buf)
		if err != nil {
			panic(err)
		}
	}
}

// Packet returns builded packet
func (p *Builder) Packet(id int32) Packet {
	return Packet{ID: id, Data: p.buf.Bytes()}
}
