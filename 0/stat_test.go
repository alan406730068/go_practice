package stat

import "testing"

func TestMean1(t *testing.T) {
	if Mean([]float64{1, 2, 3}) != 2 {
		t.Error("fail")
	}
}

func TestMean2(t *testing.T) {
	if Mean([]float64{1, 9, 5}) != 5 {
		t.Error("fail")
	}
}

func TestMean3(t *testing.T) {
	if Mean([]float64{6, 7, 10}) != 23.0/3.0 {
		t.Error("fail")
	}
}
