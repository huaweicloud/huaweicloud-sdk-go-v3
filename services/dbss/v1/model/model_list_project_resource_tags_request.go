package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListProjectResourceTagsRequest struct {

	// 资源类型。审计：auditInstance
	ResourceType string `json:"resource_type"`
}

func (o ListProjectResourceTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectResourceTagsRequest struct{}"
	}

	return strings.Join([]string{"ListProjectResourceTagsRequest", string(data)}, " ")
}
