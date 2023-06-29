package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListImageTagsRequest Request Object
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
