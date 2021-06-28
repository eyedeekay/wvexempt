package checkexempt

import "testing"

func TestGetExemptions(t *testing.T) {
	out, err := GetExemptions()
	if err != nil {
		t.Error(err)
	}
	t.Log(out)
}

func TestHasExemption(t *testing.T) {
	if ok := CheckExempt(); ok {
		t.Log("Check is:", ok)
	} else {
		t.Log("Check is:", ok)
	}
}

func TestElevateExemption(t *testing.T) {
	if err := Exempt(); err != nil {
		t.Error(err)
	}
}
