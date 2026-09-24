package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInternalEndpointPermissionsRequest Request Object
type ListInternalEndpointPermissionsRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`

	// 返回条数，默认为10，最大值为500。**注意：offset和limit参数需要配套使用。**
	Limit *int32 `json:"limit,omitempty"`

	// 起始索引，默认为0。**注意：offset和limit参数需要配套使用。**
	Offset *int32 `json:"offset,omitempty"`

	// 权限ID，用于过滤白名单权限，格式为“iam:domain::domain_id”。 其中\"domain_id\"为授权用户的账号ID， 例如“iam:domain::6e9dfd51d1124e8d8498dce894923a0d”，支持模糊搜索。
	Permission *string `json:"permission,omitempty"`
}

func (o ListInternalEndpointPermissionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInternalEndpointPermissionsRequest struct{}"
	}

	return strings.Join([]string{"ListInternalEndpointPermissionsRequest", string(data)}, " ")
}
