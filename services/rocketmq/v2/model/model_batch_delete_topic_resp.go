package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteTopicResp struct {

	// 删除主题任务ID。
	JobId *string `json:"job_id,omitempty"`
}

func (o BatchDeleteTopicResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteTopicResp struct{}"
	}

	return strings.Join([]string{"BatchDeleteTopicResp", string(data)}, " ")
}
