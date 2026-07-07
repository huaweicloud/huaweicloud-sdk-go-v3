package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StartAnalysisSessionRequestBody struct {

	// 收集会话请求ID
	CollectRequestId string `json:"collect_request_id"`
}

func (o StartAnalysisSessionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartAnalysisSessionRequestBody struct{}"
	}

	return strings.Join([]string{"StartAnalysisSessionRequestBody", string(data)}, " ")
}
