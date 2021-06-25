package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateAssetResponse struct {
	// 媒体ID<br/>

	AssetId *string `json:"asset_id,omitempty"`
	// 原始视频文件的下载url<br/>

	VideoUploadUrl *string `json:"video_upload_url,omitempty"`
	// 封面文件（Cover0）的下载url<br/>

	CoverUploadUrl *string `json:"cover_upload_url,omitempty"`
	// 字幕文件上传url数组<br/>

	SubtitleUploadUrls *[]string `json:"subtitle_upload_urls,omitempty"`
	HttpStatusCode     int       `json:"-"`
}

func (o UpdateAssetResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateAssetResponse struct{}"
	}

	return strings.Join([]string{"UpdateAssetResponse", string(data)}, " ")
}
