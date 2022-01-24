package main

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

func passOrFail() (string, error) {
	r := bufio.NewReader(os.Stdin)
	v, err := r.ReadString('\n')
	v = strings.TrimSpace(v)
	grade, err2 := strconv.Atoi(v)
	if err2 != nil {
		return v, errors.New("converting error")
	}

	if err != nil {
		return v, errors.New("reading error")
	}
	if grade < 50 {
		return "kaldi", nil
	}
	return "gecti", nil
}
