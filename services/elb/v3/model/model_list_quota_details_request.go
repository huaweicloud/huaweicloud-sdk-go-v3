package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListQuotaDetailsRequest struct {

	// 资源类型。  取值： loadbalancer、listener、ipgroup、pool、member、members_per_pool、 healthmonitor、l7policy、certificate、security_policy、 ipgroup_bindings、ipgroup_max_length。  members_per_pool表示一个pool下最多可关联的member数量。  ipgroup_bindings表示一个ipgroup下最多可关联的listener数量。  ipgroup_max_length表示一个ipgroup下最多设置的ip地址数量。  支持多值查询，查询条件格式：quota_key=xxx&quota_key=xxx。
	QuotaKey *[]string `json:"quota_key,omitempty"`
}

func (o ListQuotaDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQuotaDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListQuotaDetailsRequest", string(data)}, " ")
}
