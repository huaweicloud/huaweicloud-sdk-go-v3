package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListQuotaDetailsRequest Request Object
type ListQuotaDetailsRequest struct {

	// 资源类型。  取值： loadbalancer、listener、ipgroup、pool、member、healthmonitor、l7policy、certificate、security_policy、listeners_per_loadbalancer、listeners_per_pool、members_per_pool、condition_per_policy、ipgroup_bindings、ipgroup_max_length。  支持多值查询，查询条件格式：quota_key=xxx&quota_key=xxx。
	QuotaKey *[]string `json:"quota_key,omitempty"`
}

func (o ListQuotaDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQuotaDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListQuotaDetailsRequest", string(data)}, " ")
}
