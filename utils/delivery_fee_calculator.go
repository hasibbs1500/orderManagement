package utils

func CalculateDeliveryFee(cityID int, weight float64) float64 {
	var basePrice float64

	if cityID == 1 {
		if weight <= 0.5 {
			basePrice = 60
		} else if weight <= 1 {
			basePrice = 70
		} else {
			basePrice = 70 + ((weight - 1) * 15)
		}
	} else {
		basePrice = 100
		basePrice += (weight - 0.5) * 15
	}

	return basePrice
}
