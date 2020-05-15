package main

func checkScoreAverage(score float64) string {
	var v = int(score)
	switch {
	case v >= 9 && v <= 10:
		return "A"
	case v >= 8 && v < 9:
		return "B"
	case v >= 7 && v < 8:
		return "C"
	case v >= 6 && v < 7:
		return "D"
	case v >= 5 && v < 6:
		return "E"
	case v >= 4 && v < 5:
		return "F"
	case v >= 0 && v < 4:
		return "G"
	default:
		return "invalid average"
	}
}
