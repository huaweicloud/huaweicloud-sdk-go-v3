package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueryAnalysisResultBody 获取分析结果请求体
type QueryAnalysisResultBody struct {

	// 数据库引擎类型
	DatastoreType string `json:"datastore_type"`

	// 节点ID
	NodeId *string `json:"node_id,omitempty"`

	// 开始时间（Unix timestamp），单位：毫秒
	StartTime int64 `json:"start_time"`

	// 结束时间（Unix timestamp），单位：毫秒
	EndTime int64 `json:"end_time"`

	// CES指标
	Metrics *[]string `json:"metrics,omitempty"`

	// 动作（relatedSqlAnalysis/slowSqlAnalysis）
	Action string `json:"action"`
}

func (o QueryAnalysisResultBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryAnalysisResultBody struct{}"
	}

	return strings.Join([]string{"QueryAnalysisResultBody", string(data)}, " ")
}
