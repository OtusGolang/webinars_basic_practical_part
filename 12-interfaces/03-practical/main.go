package main

import (
	"webinars_basic_practical_part/12-interfaces/03-practical/bl"
	"webinars_basic_practical_part/12-interfaces/03-practical/repo"
)

func main() {
	realRepo := repo.NewRealDbRepo()
	bl.DoBigOperation(realRepo)
}
