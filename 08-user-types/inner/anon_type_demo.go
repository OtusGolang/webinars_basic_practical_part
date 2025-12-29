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
