package service

import (
	"errors"
	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
	"strings"
)

func Convert(input string) (string, error) {
	s := strings.TrimSpace(input)

	if s == "" {
		return "", errors.New("input is empty")
	}

	isMorse := true
	for _, r := range s {
		if r != '.' && r != '-' && r != ' ' {
			isMorse = false
			break
		}
	}
	if isMorse {
		return morse.ToText(input), nil
	}
	return morse.ToMorse(input), nil
}
