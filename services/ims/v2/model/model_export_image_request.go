package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ExportImageRequest struct {
	// 镜像ID。

	ImageId string `json:"image_id"`

	Body *ExportImageRequestBody `json:"body,omitempty"`
}

func (o ExportImageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportImageRequest struct{}"
	}

	return strings.Join([]string{"ExportImageRequest", string(data)}, " ")
}
