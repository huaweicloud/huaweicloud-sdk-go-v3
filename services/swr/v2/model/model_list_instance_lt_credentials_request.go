package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceLtCredentialsRequest Request Object
type ListInstanceLtCredentialsRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`

	// 起始索引，默认为0。**注意：offset和limit参数需要配套使用。**
	Offset *int32 `json:"offset,omitempty"`

	// 返回条数，默认为100，最大值为100。**注意：offset和limit参数需要配套使用。**
	Limit *int32 `json:"limit,omitempty"`

	// 值为false的时候，拥有te_admin角色的用户可以查询实例下所有的长期登录凭证，默认情况下只查询自己创建的长期登录凭证
	SelfOnly *bool `json:"self_only,omitempty"`
}

func (o ListInstanceLtCredentialsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceLtCredentialsRequest struct{}"
	}

	return strings.Join([]string{"ListInstanceLtCredentialsRequest", string(data)}, " ")
}
