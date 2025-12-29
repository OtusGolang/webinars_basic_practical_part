package inner

type Animal struct {
	Name string
}
type Creature struct {
	Something string
}

type Dog struct {
	Animal
	Creature
	TailLength int
}
