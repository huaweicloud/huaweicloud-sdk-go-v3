package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListImageTagsRequest struct {
	// 镜像ID。

	ImageId string `json:"image_id"`
}

func (o ListImageTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImageTagsRequest struct{}"
	}

	return strings.Join([]string{"ListImageTagsRequest", string(data)}, " ")
}
