package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateAssetByFileUploadResponse struct {

	// 媒体ID
	AssetId *string `json:"asset_id,omitempty" xml:"asset_id"`

	// 视频上传URL
	VideoUploadUrl *string `json:"video_upload_url,omitempty" xml:"video_upload_url"`

	// 封面上传地址
	CoverUploadUrl *string `json:"cover_upload_url,omitempty" xml:"cover_upload_url"`

	// 字幕文件上传url数组
	SubtitleUploadUrls *[]string `json:"subtitle_upload_urls,omitempty" xml:"subtitle_upload_urls"`

	Target         *FileAddr `json:"target,omitempty" xml:"target"`
	HttpStatusCode int       `json:"-"`
}

func (o CreateAssetByFileUploadResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAssetByFileUploadResponse struct{}"
	}

	return strings.Join([]string{"CreateAssetByFileUploadResponse", string(data)}, " ")
}
