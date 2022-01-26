package myPackage_test

import (
	"fmt"
	"golangTutorial/myPackage"
	"testing"
)

//?? Testing first aproach.
// first aproach is inefficient..
func TestDivide(t *testing.T) {
	_, err := myPackage.Divide(1.4, 1.0)
	if err != nil {
		fmt.Println("error dividing by 1.4 to 1.0")
		return
	}
	fmt.Println("success!")
}

func TestBadDivide(t *testing.T) {
	_, err := myPackage.Divide(10.0, 0.0)
	if err == nil {
		fmt.Println("success!")
		return
	}
	fmt.Println("there is an err dividing by 10.0 to 0.0")
} //
//end of first aproach.

// to use second aproach test method.
// testing data.
var tests = []struct {
	name     string
	divident float64
	divisor  float64
	expected float64
	isErr    bool
}{
	{"valid data", 6.3, 2.1, 3.0, false},
	{"dividing by zero error", 5.1, 0.0, 0, true},
}

// SECOND APROACH
// this is better than first one.
func TestDivision(t *testing.T) {
	for _, tt := range tests {
		got, err := myPackage.Divide(tt.divident, tt.divisor)
		if tt.isErr {
			if err == nil {
				t.Error("expected an error but did not get one")
			}
		} else {
			if err != nil {
				t.Error("did not expected an error but got one", err.Error())
			}
		}

		if got != tt.expected {
			t.Errorf("expected %f but got %f", tt.expected, got)
		}
	}
}

/*
call quote.Go()
crete packageName_test.go file
create function which starting with "Test<something>" (etc: TestMyPackage(t *testing.T))
to write test result t.Fatal("write something")
*/

/*
to run second aproach
->cd myPackage
-> go test -v
then you will see the tests detail.

-> go test cover => to see that how much tests succed as a percentage
-> go test -coverprofile=coverage.out && go tool cover -html=coverage.out => to see which one coverd or not in the html file.
*/
