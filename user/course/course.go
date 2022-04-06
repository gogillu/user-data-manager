package course

type Course string

const (
	A Course = "A"
	B Course = "B"
	C Course = "C"
	D Course = "D"
	E Course = "E"
	F Course = "F"
)

var AllCourses = map[string]Course{
	"A": A,
	"B": B,
	"C": C,
	"D": D,
	"E": E,
	"F": F,
}
