package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ServerJobDeleteRequest struct {

	// **参数解释**：要删除的DevServer的任务id列表
	JobIds *[]string `json:"job_ids,omitempty"`
}

func (o ServerJobDeleteRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerJobDeleteRequest struct{}"
	}

	return strings.Join([]string{"ServerJobDeleteRequest", string(data)}, " ")
}
