package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPermissionsRequest Request Object
type ListPermissionsRequest struct {

	// 分页页面的最大值。
	Limit *int32 `json:"limit,omitempty"`

	// 页面标记。
	Marker *string `json:"marker,omitempty"`

	// 资源类型的名称。
	ResourceType *string `json:"resource_type,omitempty"`
}

func (o ListPermissionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPermissionsRequest struct{}"
	}

	return strings.Join([]string{"ListPermissionsRequest", string(data)}, " ")
}
