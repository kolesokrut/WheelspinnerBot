package youtube

type YouTubeMP3 struct {
	Link     string  `json:"link"`
	Title    string  `json:"title"`
	Progress int     `json:"progress"`
	Duration float64 `json:"duration"`
	Status   string  `json:"status"`
	Msg      string  `json:"msg"`
}
