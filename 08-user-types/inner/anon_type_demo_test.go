package inner

type Animal struct {
	Name  string
	Color string
}

type Pet struct {
	Owner    string
	Nickname string
}

type Dog struct {
	Animal
	Pet
	TailLength int
}
