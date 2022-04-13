package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 请求参数
type AddOrUpdateTagsRequestBody struct {
	// 镜像ID。

	ImageId string `json:"image_id"`
	// 标签数据。 tag和image_tag只能使用一个。

	Tag *string `json:"tag,omitempty"`

	ImageTag *ResourceTag `json:"image_tag,omitempty"`
}

func (o AddOrUpdateTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddOrUpdateTagsRequestBody struct{}"
	}

	return strings.Join([]string{"AddOrUpdateTagsRequestBody", string(data)}, " ")
}
