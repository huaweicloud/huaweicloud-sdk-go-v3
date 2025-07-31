package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CopyImageInRegionInSafeModeRequest Request Object
type CopyImageInRegionInSafeModeRequest struct {

	// 镜像ID
	ImageId string `json:"image_id"`

	Body *CopyImageInRegionRequestBody `json:"body,omitempty"`
}

func (o CopyImageInRegionInSafeModeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyImageInRegionInSafeModeRequest struct{}"
	}

	return strings.Join([]string{"CopyImageInRegionInSafeModeRequest", string(data)}, " ")
}
