package youtube

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func getDownloadURL(videoURL, api string) (link string, err error) {
	url := fmt.Sprintf("https://youtube-mp36.p.rapidapi.com/dl?id=%s", videoURL)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", api)
	req.Header.Add("X-RapidAPI-Host", "youtube-mp36.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var result YouTubeMP3
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		return "", err
	}

	return result.Link, nil
}

func DownloadMP3(uri, api string) (string, error) {
	u, err := url.Parse(uri)
	if err != nil {
		log.Fatal(err)
	}

	q := u.Query()
	k := u.Path[1:]
	url := q["v"]

	id, err := getDownloadURL(strings.Join(url, " "), api)
	if err != nil {
		return "", err
	}

	if len(url) == 0 {
		if id, err = getDownloadURL(k, api); err != nil {
			return "", err
		}
	}
	return id, nil
}
