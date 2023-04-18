package helper

import (
	"io/ioutil"
	"net/http"
)

func CallApi(url string) ([]byte, error) {
	resp, err := http.Get(url)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
