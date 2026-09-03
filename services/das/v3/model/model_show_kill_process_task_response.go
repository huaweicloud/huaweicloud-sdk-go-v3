package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowKillProcessTaskResponse Response Object
type ShowKillProcessTaskResponse struct {

	// 自动kill会话任务列表
	SqlKillingTaskRespList *[]SqlKillingTaskResp `json:"sql_killing_task_resp_list,omitempty"`
	HttpStatusCode         int                   `json:"-"`
}

func (o ShowKillProcessTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKillProcessTaskResponse struct{}"
	}

	return strings.Join([]string{"ShowKillProcessTaskResponse", string(data)}, " ")
}
