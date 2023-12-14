package net
import( 
    "io"
    "net/http"
    "log"
)

func Get(url string) ([]uint8, error) {
	response, err := http.Get(url)

    if err != nil {
        log.Fatal(err)
        return nil, err
    }

    responseData, err := io.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
        return nil, err
    }

	return responseData, nil
}