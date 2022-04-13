package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListQuotaDetailsRequest struct {
	// 资源类型，取值：loadbalancer、listener、ipgroup、pool、member、members_per_pool、healthmonitor、l7policy、certificate、security_policy，其中members_per_pool表示一个pool下最多可关联的member数量。

	Type *string `json:"type,omitempty"`
}

func (o ListQuotaDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQuotaDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListQuotaDetailsRequest", string(data)}, " ")
}
