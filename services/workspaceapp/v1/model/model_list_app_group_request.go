package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAppGroupRequest Request Object
type ListAppGroupRequest struct {

	// 单次查询的大小[1-100]
	Limit int32 `json:"limit"`

	// 查询的偏移量
	Offset int32 `json:"offset"`

	// 应用服务器组ID
	AppServerGroupId *string `json:"app_server_group_id,omitempty"`

	// 应用组ID
	AppGroupId *string `json:"app_group_id,omitempty"`

	// 应用组名称
	Name *string `json:"name,omitempty"`

	// 授权类型(APP、APP_GROUP)
	AuthorizationType *string `json:"authorization_type,omitempty"`
}

func (o ListAppGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppGroupRequest struct{}"
	}

	return strings.Join([]string{"ListAppGroupRequest", string(data)}, " ")
}
