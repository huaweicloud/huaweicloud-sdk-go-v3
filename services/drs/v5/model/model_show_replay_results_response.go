package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowReplayResultsResponse Response Object
type ShowReplayResultsResponse struct {

	// 数据总量
	TotalCount *int64 `json:"total_count,omitempty"`

	// 回放基于时间统计详细信息列表，在type为shard_statistics时返回
	ShardStatics *[]ReplayShardStaticsResp `json:"shard_statics,omitempty"`

	// 慢SQL信息列表，在type为slow_sql时返回
	SlowSqls *[]ReplaySlowSqlResp `json:"slow_sqls,omitempty"`

	// 慢SQL统计信息列表，在type为slow_sql_template时返回
	SlowSqlTemplates *[]ReplaySlowSqlTemplateResp `json:"slow_sql_templates,omitempty"`

	// 异常SQL信息列表，在type为error_sql时返回
	ErrorSqls *[]ReplayErrorSqlResp `json:"error_sqls,omitempty"`

	// 异常SQL统计信息列表，在type为error_sql_template时返回
	ErrorSqlTemplates *[]ReplayErrorSqlTemplateResp `json:"error_sql_templates,omitempty"`

	// 正在回放SQL信息列表，在type为replaying_sql时返回
	ReplayingSqls *[]ReplayingSqlResp `json:"replaying_sqls,omitempty"`

	// 回放异常SQL分类信息，在type为error_classification时返回
	ErrorClassifications *[]ReplayErrorClassification `json:"error_classifications,omitempty"`
	HttpStatusCode       int                          `json:"-"`
}

func (o ShowReplayResultsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReplayResultsResponse struct{}"
	}

	return strings.Join([]string{"ShowReplayResultsResponse", string(data)}, " ")
}
