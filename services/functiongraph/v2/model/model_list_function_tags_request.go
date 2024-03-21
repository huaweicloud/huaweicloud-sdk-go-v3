package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFunctionTagsRequest Request Object
type ListFunctionTagsRequest struct {

	// 资源类型，此处请填写functions
	ResourceType string `json:"resource_type"`

	// 资源ID
	ResourceId string `json:"resource_id"`

	// 消息体的类型（格式）
	ContentType string `json:"Content-Type"`
}

func (o ListFunctionTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFunctionTagsRequest struct{}"
	}

	return strings.Join([]string{"ListFunctionTagsRequest", string(data)}, " ")
}
