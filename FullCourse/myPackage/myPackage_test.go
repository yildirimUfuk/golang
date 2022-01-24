package myPackage_test

import (
	"golangTutorial/myPackage"
	"testing"
)

func TestMyPackage(t *testing.T) {
	if myPackage.MyPackageFunc() != "Go Gopher" {
		t.Fatal("test result: string is not same")
	}
}

/*
call quote.Go()
crete packageName_test.go file
create function which starting with "Test<something>" (etc: TestMyPackage(t *testing.T))
to write test result t.Fatal("write something")
*/
