package enum

//go:generate enumer -type=Course

type Course int

const (
	A Course = iota
	B
	C
	D
	E
	F
)

func String(course Course) string {
	return course.String()
}
