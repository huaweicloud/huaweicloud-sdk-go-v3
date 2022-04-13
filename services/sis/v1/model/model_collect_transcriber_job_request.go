package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CollectTranscriberJobRequest struct {
	// 录音文件识别任务标识符。

	JobId string `json:"job_id"`
}

func (o CollectTranscriberJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CollectTranscriberJobRequest struct{}"
	}

	return strings.Join([]string{"CollectTranscriberJobRequest", string(data)}, " ")
}
