package tiktok

type TikTok struct {
	Video                    []string `json:"video"`
	Music                    []string `json:"music"`
	Cover                    []string `json:"cover"`
	OriginalWatermarkedVideo []string `json:"OriginalWatermarkedVideo"`
	DynamicCover             []string `json:"dynamic_cover"`
	Author                   []string `json:"author"`
	Region                   []string `json:"region"`
	AvatarThumb              []string `json:"avatar_thumb"`
	CustomVerify             []string `json:"custom_verify"`
}
