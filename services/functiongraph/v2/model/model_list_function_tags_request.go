package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFunctionTagsRequest Request Object
type ListFunctionTagsRequest struct {

	// 资源类型
	ResourceType string `json:"resource_type"`

	// 资源ID
	ResourceId string `json:"resource_id"`
}

func (o ListFunctionTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFunctionTagsRequest struct{}"
	}

	return strings.Join([]string{"ListFunctionTagsRequest", string(data)}, " ")
}
