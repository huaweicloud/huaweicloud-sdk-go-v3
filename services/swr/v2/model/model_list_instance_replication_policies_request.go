package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceReplicationPoliciesRequest Request Object
type ListInstanceReplicationPoliciesRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`

	// 排序字段，支持created_at、updated_at、name，默认为created_at
	OrderColumn *string `json:"order_column,omitempty"`

	// 排序方式，支持desc、asc，默认desc; 注意：order_type需要与order_column配合使用。
	OrderType *string `json:"order_type,omitempty"`

	// 名称，模糊查询
	Name *string `json:"name,omitempty"`

	// 仓库ID
	RegistryId *int32 `json:"registry_id,omitempty"`

	// 起始索引，默认值为0。**注意：offset和limit参数需要配套使用，offset必须为0或者为limit的倍数。**
	Offset *int32 `json:"offset,omitempty"`

	// 返回条数，默认为10，最大值为100。**注意：offset和limit参数需要配套使用，offset必须为0或者为limit的倍数。**
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListInstanceReplicationPoliciesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceReplicationPoliciesRequest struct{}"
	}

	return strings.Join([]string{"ListInstanceReplicationPoliciesRequest", string(data)}, " ")
}
