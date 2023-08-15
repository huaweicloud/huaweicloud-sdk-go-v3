package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAclPolicyBindedToApiV2Request Request Object
type ListAclPolicyBindedToApiV2Request struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	// 偏移量，表示从此偏移量开始查询，偏移量小于0时，自动转换为0
	Offset *int64 `json:"offset,omitempty"`

	// 每页显示的条目数量，条目数量小于等于0时，自动转换为20，条目数量大于500时，自动转换为500
	Limit *int32 `json:"limit,omitempty"`

	// API编号
	ApiId string `json:"api_id"`

	// 环境编号
	EnvId *string `json:"env_id,omitempty"`

	// 环境名称
	EnvName *string `json:"env_name,omitempty"`

	// ACL策略编号
	AclId *string `json:"acl_id,omitempty"`

	// ACL策略名称
	AclName *string `json:"acl_name,omitempty"`
}

func (o ListAclPolicyBindedToApiV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAclPolicyBindedToApiV2Request struct{}"
	}

	return strings.Join([]string{"ListAclPolicyBindedToApiV2Request", string(data)}, " ")
}
