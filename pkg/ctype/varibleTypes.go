package ctype

import "strconv"

func ToInt(value interface{}) int {
	// Tür kontrolü ve dönüştürme
	if value == nil {
		return 0
	}

	switch v := value.(type) {
	case int:
		return v
	case int8:
		return int(v)
	case int16:
		return int(v)
	case int32:
		return int(v)
	case int64:
		return int(v)
	case float32:
		return int(v)
	case float64:
		return int(v)
	case byte:
		return int(v)
	case string:
		parsed, err := strconv.Atoi(v)
		if err == nil {
			return parsed
		}
	}
	return 0 // Varsayılan değer, dönüşüm başarısız olursa
}

// ToBool converts different types to bool, returns false for nil, 0, or equivalent
func ToBool(value interface{}) bool {
	switch v := value.(type) {
	case nil:
		return false
	case bool:
		return v
	case int, int32, int64:
		return v != 0
	case byte, float32, float64:
		return v != 0.0
	case string:
		if len(v) > 1 {
			return true
		}
		return false
	default:
		return true
	}
}
