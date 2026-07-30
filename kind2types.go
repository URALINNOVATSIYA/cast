package cast

import (
	"reflect"
	"time"

	"github.com/google/uuid"
)

var kind2types = []reflect.Type{
	nil,
	reflect.TypeFor[bool](),
	reflect.TypeFor[int](),
	reflect.TypeFor[int8](),
	reflect.TypeFor[int16](),
	reflect.TypeFor[int32](),
	reflect.TypeFor[int64](),
	reflect.TypeFor[uint](),
	reflect.TypeFor[uint8](),
	reflect.TypeFor[uint16](),
	reflect.TypeFor[uint32](),
	reflect.TypeFor[uint64](),
	reflect.TypeFor[uintptr](),
	reflect.TypeFor[float32](),
	reflect.TypeFor[float64](),
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	reflect.TypeFor[string](),
	nil,
	nil,
}

var typeByteSlice = reflect.TypeFor[[]byte]()
var typeUUID = reflect.TypeFor[uuid.UUID]()
var typeTime = reflect.TypeFor[time.Time]()
