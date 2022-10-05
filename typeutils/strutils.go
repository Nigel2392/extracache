package typeutils

type CanContain interface {
	int | string | rune | byte | float32 | float64 | bool | complex64 | complex128 | int8 | int16 | int64 | uint | uint16 | uint32 | uint64 | uintptr
}

func Contains[T CanContain](arr []T, item T) bool {
	for _, v := range arr {
		if v == item {
			return true
		}
	}
	return false
}

func PadLeft(str string, length int, pad string) string {
	if len(str) >= length {
		return str
	}
	return PadLeft(pad+str, length, pad)
}
func PadRight(str string, length int, pad string) string {
	if len(str) >= length {
		return str
	}
	return PadRight(str+pad, length, pad)
}

func Repeat(str string, count int) string {
	if count <= 0 {
		return ""
	}
	return str + Repeat(str, count-1)
}
