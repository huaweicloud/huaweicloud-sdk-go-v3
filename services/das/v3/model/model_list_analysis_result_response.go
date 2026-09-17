package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAnalysisResultResponse Response Object
type ListAnalysisResultResponse struct {

	// 是否成功
	Success *bool `json:"success,omitempty"`

	// 是否有风险
	Risk *bool `json:"risk,omitempty"`

	// 分析数据
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListAnalysisResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAnalysisResultResponse struct{}"
	}

	return strings.Join([]string{"ListAnalysisResultResponse", string(data)}, " ")
}
