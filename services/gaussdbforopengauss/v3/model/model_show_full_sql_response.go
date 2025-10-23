package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFullSqlResponse Response Object
type ShowFullSqlResponse struct {

	// **参数解释**: 链路详情。
	TraceStatistics *interface{} `json:"trace_statistics,omitempty"`

	// **参数解释**: 组件上SQL执行记录列表。
	Components     *[]FullSqlComponetResult `json:"components,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ShowFullSqlResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFullSqlResponse struct{}"
	}

	return strings.Join([]string{"ShowFullSqlResponse", string(data)}, " ")
}
