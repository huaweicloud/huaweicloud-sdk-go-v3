package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAnalysisScriptsResponse Response Object
type ListAnalysisScriptsResponse struct {

	// 计数
	Count *int64 `json:"count,omitempty"`

	// 分析脚本
	Records        *[]AnalysisScript `json:"records,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListAnalysisScriptsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAnalysisScriptsResponse struct{}"
	}

	return strings.Join([]string{"ListAnalysisScriptsResponse", string(data)}, " ")
}
