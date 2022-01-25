package myPackage

import "errors"

func Divide(x, y float64) (float64, error) {
	var result float64
	if y == 0 {
		return result, errors.New("connot divide by 0")
	}
	result = x / y
	return result, nil
}

/*
go mod init "name"
create a folder for package.
create .go file by same name of package name.
*/
