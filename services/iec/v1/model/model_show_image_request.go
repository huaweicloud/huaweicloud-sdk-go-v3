package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowImageRequest struct {
	// 镜像ID。

	ImageId string `json:"image_id"`
}

func (o ShowImageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowImageRequest struct{}"
	}

	return strings.Join([]string{"ShowImageRequest", string(data)}, " ")
}
