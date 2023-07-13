package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type JobIdInfo struct {

	// 任务ID
	JobId *string `json:"job_id,omitempty"`
}

func (o JobIdInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobIdInfo struct{}"
	}

	return strings.Join([]string{"JobIdInfo", string(data)}, " ")
}
