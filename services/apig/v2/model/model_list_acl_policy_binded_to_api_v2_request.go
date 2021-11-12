package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListAclPolicyBindedToApiV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`
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
	// 偏移量，表示从此偏移量开始查询，偏移量小于0时，自动转换为0

	Offset *int64 `json:"offset,omitempty"`
	// 每页显示的条目数量

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListAclPolicyBindedToApiV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAclPolicyBindedToApiV2Request struct{}"
	}

	return strings.Join([]string{"ListAclPolicyBindedToApiV2Request", string(data)}, " ")
}
