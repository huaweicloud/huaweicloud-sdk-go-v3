package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateAssetByFileUploadResponse struct {
	// 媒体ID<br/>

	AssetId *string `json:"asset_id,omitempty"`
	// 视频上传URL<br/>

	VideoUploadUrl *string `json:"video_upload_url,omitempty"`
	// 封面上传地址<br/>

	CoverUploadUrl *string `json:"cover_upload_url,omitempty"`
	// 字幕文件上传url数组<br/>

	SubtitleUploadUrls *[]string `json:"subtitle_upload_urls,omitempty"`

	Target         *FileAddr `json:"target,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o CreateAssetByFileUploadResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateAssetByFileUploadResponse struct{}"
	}

	return strings.Join([]string{"CreateAssetByFileUploadResponse", string(data)}, " ")
}
