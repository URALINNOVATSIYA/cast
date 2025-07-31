package cast

import (
	"testing"

	"github.com/google/uuid"
)

func TestAsUuid(t *testing.T) {
	id := uuid.New()
	strId := id.String()
	binaryId, _ := id.MarshalBinary()
	byteId, _ := id.MarshalText()
	tests := []castTest[uuid.UUID]{
		{nil, uuid.Nil, ""},
		{strId, id, ""},
		{binaryId, id, ""},
		{byteId, id, ""},
		{"", uuid.Nil, "invalid UUID length: 0"},
		{"c3a120f3-a594-4ec-b7f1-5e63f5bf834d", uuid.Nil, "invalid UUID length: 35"},
		{123, uuid.Nil, "failed to cast int to UUID"},
	}
	runCastTests(t, "AsUuid", AsUuid, tests)
}
