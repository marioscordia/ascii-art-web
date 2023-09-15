package mainfunc

import (
	"errors"

	"ascii-art-web/funcs/checker"
	"ascii-art-web/funcs/internal"
)

func WriteStr(str, banner string) (string, error) {
	err := checker.CheckWord(str)
	if err != nil {
		return "", errors.New("error")
	}
	hash, err := internal.CreateHash(banner)
	if err != nil {
		return "", errors.New("error")
	}
	res := internal.Write(str, hash)
	return res, nil
}
