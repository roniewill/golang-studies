package main

func scoreAverage(score float64) string {
	/*
		WARNING!
		switch clause does not automatically pass to the next case,
		on the contrary it is necessary to use fallthrough for this to happen.

		A clausula switch não passa automaticamente para o próximo case, do contrário,
		é necessário usar fallthrough pra isso acontencer.
	*/

	var average = int(score)
	switch average {
	case 10:
		fallthrough
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "invalid average"
	}
}
