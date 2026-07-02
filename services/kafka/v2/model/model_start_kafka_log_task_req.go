package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StartKafkaLogTaskReq struct {

	// **参数解释**： 日志任务列表。
	LogTaskList *[]StartKafkaLogTaskReqEntity `json:"log_task_list,omitempty"`
}

func (o StartKafkaLogTaskReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartKafkaLogTaskReq struct{}"
	}

	return strings.Join([]string{"StartKafkaLogTaskReq", string(data)}, " ")
}
