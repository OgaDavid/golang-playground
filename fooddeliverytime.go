package main

type food struct {
	name     string
	prepTime int
}

func FoodDeliveryTime(order string) int {
	var burger = food{
		name:     "burger",
		prepTime: 15,
	}

	var chips = food{
		name:     "chips",
		prepTime: 10,
	}

	var nuggets = food{
		name:     "nuggets",
		prepTime: 12,
	}

	menu := []food{
		burger,
		chips,
		nuggets,
	}

	for _, food := range menu {
		if order == food.name {
			return food.prepTime
		}
	}

	return 404
}
