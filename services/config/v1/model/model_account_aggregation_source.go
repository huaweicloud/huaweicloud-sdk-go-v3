package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AccountAggregationSource 聚合数据的源帐号。
type AccountAggregationSource struct {

	// 帐号列表。
	DomainIds *[]string `json:"domain_ids,omitempty"`
}

func (o AccountAggregationSource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccountAggregationSource struct{}"
	}

	return strings.Join([]string{"AccountAggregationSource", string(data)}, " ")
}
