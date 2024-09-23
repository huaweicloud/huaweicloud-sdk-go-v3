package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEntitiesRequest Request Object
type ListEntitiesRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	// 父节点（根或组织单元）的唯一标识符（ID）。
	ParentId *string `json:"parent_id,omitempty"`

	// 子节点（组织单元）的唯一标识符（ID）。
	ChildId *string `json:"child_id,omitempty"`

	// 页面中最大结果数量。
	Limit *int32 `json:"limit,omitempty"`

	// 分页标记。
	Marker *string `json:"marker,omitempty"`
}

func (o ListEntitiesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEntitiesRequest struct{}"
	}

	return strings.Join([]string{"ListEntitiesRequest", string(data)}, " ")
}
