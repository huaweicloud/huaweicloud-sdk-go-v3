package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowProjectTagsListRequest Request Object
type ShowProjectTagsListRequest struct {

	// 资源类型，此处请填写functions
	ResourceType string `json:"resource_type"`

	// 消息体的类型（格式）
	ContentType string `json:"Content-Type"`
}

func (o ShowProjectTagsListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProjectTagsListRequest struct{}"
	}

	return strings.Join([]string{"ShowProjectTagsListRequest", string(data)}, " ")
}
