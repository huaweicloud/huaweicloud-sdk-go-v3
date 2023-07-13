package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadImageFileRequest Request Object
type DownloadImageFileRequest struct {

	// devcloud项目的32位id
	ProjectId string `json:"project_id"`

	// 图片URI
	ImageUri string `json:"image_uri"`
}

func (o DownloadImageFileRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadImageFileRequest struct{}"
	}

	return strings.Join([]string{"DownloadImageFileRequest", string(data)}, " ")
}
