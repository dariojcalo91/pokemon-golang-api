package shared

type Field string

func (f Field) String() string {
	return string(f)
}
