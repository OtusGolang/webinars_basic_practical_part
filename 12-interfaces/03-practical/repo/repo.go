package repo

import (
	"fmt"
	"webinars_basic_practical_part/12-interfaces/02-practical/models"
)

type RealDbRepo struct{}

func NewRealDbRepo() *RealDbRepo {
	return &RealDbRepo{}
}

func (r *RealDbRepo) SaveItem(item models.ItemDbModel) {
	fmt.Printf("Item saved to db: %s, Price: %.2f\n", item.Name, item.Price)
	// Here we use DB connection to save object to a real db
}
