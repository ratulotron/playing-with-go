package people

// Human is a human
type Human struct {
	name  string
	age   int
	phone string
}

// Student has Human atrributes
type Student struct {
	Human
	school string
	loan   float32
}

// Employee is a human
type Employee struct {
	Human
	company string
	money   float32
}

// Men is a interface for humans
type Men interface {
	SayHi()
	Sing(lyrics string)
	Guzzle(beerStein string)
}
