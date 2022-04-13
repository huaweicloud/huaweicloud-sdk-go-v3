package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CopyImageInRegionRequest struct {
	// 镜像ID

	ImageId string `json:"image_id"`

	Body *CopyImageInRegionRequestBody `json:"body,omitempty"`
}

func (o CopyImageInRegionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyImageInRegionRequest struct{}"
	}

	return strings.Join([]string{"CopyImageInRegionRequest", string(data)}, " ")
}
