package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CopyImageCrossRegionRequest struct {

	// 镜像ID
	ImageId string `json:"image_id"`

	Body *CopyImageCrossRegionRequestBody `json:"body,omitempty"`
}

func (o CopyImageCrossRegionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyImageCrossRegionRequest struct{}"
	}

	return strings.Join([]string{"CopyImageCrossRegionRequest", string(data)}, " ")
}
