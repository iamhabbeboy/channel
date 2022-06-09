package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func DoContext() {
	timeoutContext, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
	defer cancel()

	req, err := http.NewRequestWithContext(timeoutContext, http.MethodGet, "http://placehold.it/2000x2000", nil)

	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	imageData, err := ioutil.ReadAll(res.Body)
	fmt.Printf("downloaded image of size %d\n", len(imageData))
}
