package soundcloud

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func getLinkFromResponse(URL string) string {
	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	f := func(i int, s *goquery.Selection) bool {
		link, _ := s.Attr("rel")
		return strings.HasPrefix(link, "canonical")
	}
	var link string
	doc.Find("link").FilterFunction(f).Each(func(_ int, tag *goquery.Selection) {
		link, _ = tag.Attr("href")
	})
	return link
}

func DownloadMusic(link, api string) string {
	URL := fmt.Sprintf("https://soundcloud4.p.rapidapi.com/song/download?track_url=%s", getLinkFromResponse(link))

	req, _ := http.NewRequest("GET", URL, nil)

	req.Header.Add("X-RapidAPI-Key", api)
	req.Header.Add("X-RapidAPI-Host", "soundcloud4.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var result SoundCloud
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	return result.URL
}
