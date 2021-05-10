package server

import (
	"encoding/binary"

	"github.com/Toffee-iZt/gomine/common"
)

const (
	// TPS is ticks per second
	TPS = 20
	// MPT is milliseconds per tick
	MPT = 1000 / TPS
)

// DefaultWorldHashedSeed ...
var DefaultWorldHashedSeed = int64(binary.LittleEndian.Uint64(common.JavaSHA256HashLong(common.JavaStringHashCode("North Carolina"))))
