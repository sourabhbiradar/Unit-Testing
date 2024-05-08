package mymath

import (
	"testing"
	"time"
)

func TestAbs(t *testing.T) { // testing Abs()
	// testing with n = -2 , 0 , 2

	if Abs(-2) < 0 {
		t.Error("Negetive value found in Abs() :", -2)
	}

	if Abs(0) < 0 {
		t.Error("Negetive value found in Abs() :", 0)
	}

	if Abs(2) < 0 {
		t.Error("Negetive value found in Abs() :", 2)
	}
}

// sub-testing
func TestAbsSub(t *testing.T) {
	t.Run("negetive number", func(t *testing.T) {
		if Abs(-2) < 0 {
			t.Error("Negetive value found in Abs() :", -2)
		}
	})

	t.Run("zero", func(t *testing.T) {
		if Abs(0) < 0 {
			t.Error("Negetive value found in Abs() :", 0)
		}
	})

	t.Run("positive", func(t *testing.T) {
		if Abs(2) < 0 {
			t.Error("Negetive value found in Abs() :", 2)
		}
	})
}

// t.Skip()
func TestSkip(t *testing.T) {
	if Abs(3) == 3 {
		t.Skip("skipping becoz n can not be 3") //for ilustrastion
		t.Log("this line will be skiped")
	}

	if true {
		t.Log("just to demonstrate")
	}
}

// t.Cleanup()
func TestCleanup(t *testing.T) {
	t.Cleanup(func() {
		t.Log("clean up")
	})

	t.Log("before clean up")
}

// t.Parallel()
func TestParallel(t *testing.T) {
	// without t.Parallel() test would take >9 sec
	// with t.Parallel() 3-4 sec

	t.Run("first", func(t *testing.T) {
		t.Parallel()
		time.Sleep(3 * time.Second)
	})

	t.Run("second", func(t *testing.T) {
		t.Parallel()
		time.Sleep(3 * time.Second)
	})

	t.Run("third", func(t *testing.T) {
		t.Parallel()
		time.Sleep(3 * time.Second)
	})

}
