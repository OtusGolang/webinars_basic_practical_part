package bl

import "webinars_basic_practical_part/12-interfaces/02-practical/models"

type ItemSaver interface {
	SaveItem(item models.ItemDbModel)
}

func DoBigOperation(itemSaver ItemSaver) {
	item := models.ItemDbModel{Name: "Big Item", Price: 99.99}
	itemSaver.SaveItem(item)
	item2 := models.ItemDbModel{Name: "Another Big Item", Price: 199.99}
	itemSaver.SaveItem(item2)
	item3 := models.ItemDbModel{Name: "Third Big Item", Price: 299.99}
	itemSaver.SaveItem(item3)

}
