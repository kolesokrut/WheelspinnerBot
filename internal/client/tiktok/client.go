package tiktok

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func DownloadVideo(link string) string {
	url := fmt.Sprintf("https://tiktok-downloader-download-tiktok-videos-without-watermark.p.rapidapi.com/vid/index?url=%s", link)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "0f0ab81e36msh7cec3da23406dd7p14872ejsnb52d8ae809a8")
	req.Header.Add("X-RapidAPI-Host", "tiktok-downloader-download-tiktok-videos-without-watermark.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var result TikTok
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	justString := strings.Join(result.Video, " ")
	fmt.Println(justString)
	return justString
}
