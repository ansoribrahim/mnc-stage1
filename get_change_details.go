package main

func GetChangeDetails(total, paid int) (map[int]int, int, bool) {
	if paid < total {
		return nil, 0, false
	}

	change := paid - total
	roundedChange := (change / 100) * 100

	denominations := []int{100000, 50000, 20000, 10000, 5000, 2000, 1000, 500, 200, 100}

	result := make(map[int]int)
	remainingChange := roundedChange

	for _, denom := range denominations {
		if remainingChange >= denom {
			quantity := remainingChange / denom
			if quantity > 0 {
				result[denom] = quantity
				remainingChange = remainingChange % denom
			}
		}
	}

	return result, roundedChange, true
}
