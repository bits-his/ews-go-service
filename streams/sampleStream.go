package streams

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var (
	streamEndpoint = "https://api.twitter.com/2/tweets/sample/stream"
	bearer_token   = os.Getenv("TWITTER_BEARER_TOKEN")
	httpClient     *http.Client
	//req            *http.Request
	err error
)

func SampleStream() {

	req, err := http.NewRequest(http.MethodGet, streamEndpoint, nil)
	if err != nil {
		log.Fatalf("Encountered error forming request %v", err)
	}

	fmt.Println(bearer_token)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", bearer_token))
	req.Header.Set("User-Agent", "v2SampledStreamGo")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Encountered error making request %v", err)
	}

	defer resp.Body.Close()

	fmt.Println(resp.StatusCode)

	if resp.StatusCode == 200 {

		bodyBytes, err := io.ReadAll(resp.Body)

		// Error checking of the ioutil.ReadAll() request
		if err != nil {
			log.Fatal(err)
		}

		bodyString := string(bodyBytes)

		fmt.Println(bodyString)
	} else {
		log.Fatalf("Encountered error making request %v, %d", err, resp.StatusCode)
	}
}
