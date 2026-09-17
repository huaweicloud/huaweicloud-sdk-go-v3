package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDdsSlowLogTrendRequestBody 获取DDS慢日志趋势请求体
type ShowDdsSlowLogTrendRequestBody struct {

	// 开始时间
	StartTime int64 `json:"start_time"`

	// 结束时间
	EndTime int64 `json:"end_time"`

	// 节点ID列表
	NodeIds *[]string `json:"node_ids,omitempty"`
}

func (o ShowDdsSlowLogTrendRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDdsSlowLogTrendRequestBody struct{}"
	}

	return strings.Join([]string{"ShowDdsSlowLogTrendRequestBody", string(data)}, " ")
}
