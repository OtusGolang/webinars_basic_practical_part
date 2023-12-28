package inner

func GetInstance() struct {
	a int
	b string
} {
	return struct {
		a int
		b string
	}{99, "inner"}

}

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
