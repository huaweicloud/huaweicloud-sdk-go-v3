package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEndpointPermissionsRequest Request Object
type ListEndpointPermissionsRequest struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	// 偏移量，表示从此偏移量开始查询，偏移量小于0时，自动转换为0
	Offset *int64 `json:"offset,omitempty"`

	// 每页显示的条目数量，条目数量小于等于0时，自动转换为20，条目数量大于500时，自动转换为500
	Limit *int32 `json:"limit,omitempty"`

	// 权限帐号ID，格式为“iam:domain::domain_id”，支持模糊搜索
	Permission *string `json:"permission,omitempty"`
}

func (o ListEndpointPermissionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEndpointPermissionsRequest struct{}"
	}

	return strings.Join([]string{"ListEndpointPermissionsRequest", string(data)}, " ")
}
