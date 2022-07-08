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

func getDownloadURL(videoURL string) (link string, err error) {
	url := fmt.Sprintf("https://youtube-mp36.p.rapidapi.com/dl?id=%s", videoURL)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "0f0ab81e36msh7cec3da23406dd7p14872ejsnb52d8ae809a8")
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

func DownloadMP3(uri string) (string, error) {
	u, err := url.Parse(uri)
	if err != nil {
		log.Fatal(err)
	}

	q := u.Query()
	k := u.Path[1:]
	url := q["v"]

	id, err := getDownloadURL(strings.Join(url, " "))
	if err != nil {
		return "", err
	}

	if len(url) == 0 {
		if id, err = getDownloadURL(k); err != nil {
			return "", err
		}
	}
	return id, nil
}
