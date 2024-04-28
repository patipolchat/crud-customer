package typehelper

func GetPointer[T any](value T) *T {
	return &value
}
