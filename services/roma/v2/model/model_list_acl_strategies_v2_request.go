package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAclStrategiesV2Request Request Object
type ListAclStrategiesV2Request struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 偏移量，表示从此偏移量开始查询，偏移量小于0时，自动转换为0
	Offset *int64 `json:"offset,omitempty"`

	// 每页显示的条目数量
	Limit *int32 `json:"limit,omitempty"`

	// ACL策略编号。
	Id *string `json:"id,omitempty"`

	// ACL策略名称。
	Name *string `json:"name,omitempty"`

	// 类型 - PERMIT (白名单类型) - DENY (黑名单类型)
	AclType *string `json:"acl_type,omitempty"`

	// 作用的对象类型： - IP - DOMAIN
	EntityType *string `json:"entity_type,omitempty"`

	// 指定需要精确匹配查找的参数名称，多个参数需要支持精确匹配时参数之间使用“,”隔开。  目前仅支持name。
	PreciseSearch *string `json:"precise_search,omitempty"`
}

func (o ListAclStrategiesV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAclStrategiesV2Request struct{}"
	}

	return strings.Join([]string{"ListAclStrategiesV2Request", string(data)}, " ")
}
