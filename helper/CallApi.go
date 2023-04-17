package helper

import (
	"io/ioutil"
	"net/http"
)

func CallApi(url string) ([]byte, error) {
	// start : Call api and format data
	resp, err := http.Get(url)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
