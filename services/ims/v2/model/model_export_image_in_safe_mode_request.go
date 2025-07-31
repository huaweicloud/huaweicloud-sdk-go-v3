package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportImageInSafeModeRequest Request Object
type ExportImageInSafeModeRequest struct {

	// 镜像ID。
	ImageId string `json:"image_id"`

	Body *ExportImageRequestBody `json:"body,omitempty"`
}

func (o ExportImageInSafeModeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportImageInSafeModeRequest struct{}"
	}

	return strings.Join([]string{"ExportImageInSafeModeRequest", string(data)}, " ")
}
