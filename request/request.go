package request

import (
	"io/ioutil"
	"net/http"
)

func Curl(target string) ([]byte, error) {
	res, err := http.Get(target)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return content, nil
}
