package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteConsumerGroupResp struct {

	// 删除消费组的任务ID
	JobId *string `json:"job_id,omitempty"`
}

func (o BatchDeleteConsumerGroupResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteConsumerGroupResp struct{}"
	}

	return strings.Join([]string{"BatchDeleteConsumerGroupResp", string(data)}, " ")
}
