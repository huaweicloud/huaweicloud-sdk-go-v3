package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueryDataFilterRuleReq 查询数据过滤规则请求
type QueryDataFilterRuleReq struct {

	// 数据过滤类型，取值分别是common、config。common表示简单条件过滤， config表示关联表过滤
	FilterType string `json:"filter_type"`

	// 标记是否查询全部过滤规则（含已同步和未下发的），即 true（查询所有）、false（仅查询已同步）
	IsAllRules *bool `json:"is_all_rules,omitempty"`

	// 查询返回记录的数量限制。
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，表示查询该偏移量后面的记录。
	Offset *int32 `json:"offset,omitempty"`
}

func (o QueryDataFilterRuleReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryDataFilterRuleReq struct{}"
	}

	return strings.Join([]string{"QueryDataFilterRuleReq", string(data)}, " ")
}
