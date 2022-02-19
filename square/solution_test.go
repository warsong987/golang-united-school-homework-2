package square

import "testing"

const EXECTURE_TEXT = "The result is not equal to the expectation"

func Test_CalcTiangleSquare(t *testing.T) {
	res := CalcSquare(10.0, SidesTriangle)
	wait := float64(43.30127018922193)

	if res != wait {
		t.Fatalf(EXECTURE_TEXT)
	}
}

func Test_CalcSquareSquare(t *testing.T) {
	res := CalcSquare(10.0, SidesSquare)
	wait := float64(100)

	if res != wait {
		t.Fatalf(EXECTURE_TEXT)
	}
}

func Test_CalcSquareOval(t *testing.T) {
	res := CalcSquare(10.0, SidesCircle)
	wait := float64(7.957747154594767)

	if res != wait {
		t.Fatalf(EXECTURE_TEXT)
	}
}
