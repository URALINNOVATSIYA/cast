package cast

import (
	"fmt"

	"github.com/google/uuid"
)

func ToUuid(value any) uuid.UUID {
	r, err := AsUuid(value)
	if err != nil {
		panic(err)
	}
	return r
}

func AsUuid(value any) (uuid.UUID, error) {
	switch v := value.(type) {
	case nil:
		return uuid.Nil, nil
	case uuid.UUID:
		return v, nil
	case string:
		return uuid.Parse(v)
	case []byte:
		if len(v) == 16 {
			return uuid.FromBytes(v)
		}
		return uuid.ParseBytes(v)
	}
	return uuid.Nil, fmt.Errorf("failed to cast %T to UUID", value)
}
