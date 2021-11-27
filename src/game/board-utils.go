package game

func GetSpaces(w int) string {
	var spaces string = " "
	for i := float64(0); i < float64(float64(w / 2) * 0.7); i++ {
		spaces += " "
	}

	return spaces
}
