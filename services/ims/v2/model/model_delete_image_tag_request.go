package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteImageTagRequest struct {
	// 镜像ID。

	ImageId string `json:"image_id"`
	// 要删除的标签的键。

	Key string `json:"key"`
}

func (o DeleteImageTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteImageTagRequest struct{}"
	}

	return strings.Join([]string{"DeleteImageTagRequest", string(data)}, " ")
}
