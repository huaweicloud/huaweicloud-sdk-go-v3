package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobScriptAnalyzeRequest 脚本内容分析请求
type JobScriptAnalyzeRequest struct {

	// 脚本内容
	Content string `json:"content"`
}

func (o JobScriptAnalyzeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobScriptAnalyzeRequest struct{}"
	}

	return strings.Join([]string{"JobScriptAnalyzeRequest", string(data)}, " ")
}
