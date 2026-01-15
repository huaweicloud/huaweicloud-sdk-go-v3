package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAppAuthorizationsRequest Request Object
type ListAppAuthorizationsRequest struct {

	// 查询的偏移量，默认值0。
	Offset *int32 `json:"offset,omitempty"`

	// 单次查询的大小[1-100]，默认值10。
	Limit *int32 `json:"limit,omitempty"`

	// 应用ID。
	AppId string `json:"app_id"`

	// 用户名/用户组名。
	Name *string `json:"name,omitempty"`

	// 类型： * `SIMPLE` - 普通用户 * `USER_GROUP` - 用户组
	TargetType *string `json:"target_type,omitempty"`

	// 类型： * `ALL_USER` - 所有用户 * `ASSIGN_USER` - 部分用户
	AssignType *string `json:"assign_type,omitempty"`
}

func (o ListAppAuthorizationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppAuthorizationsRequest struct{}"
	}

	return strings.Join([]string{"ListAppAuthorizationsRequest", string(data)}, " ")
}
