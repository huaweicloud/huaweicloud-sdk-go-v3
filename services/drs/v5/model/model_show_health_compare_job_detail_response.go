package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHealthCompareJobDetailResponse Response Object
type ShowHealthCompareJobDetailResponse struct {

	// 对比任务ID。
	Id *string `json:"id,omitempty"`

	// 对比类型： object_comparison：对象对比。 lines：行对比。 account：用户对比。
	Type *string `json:"type,omitempty"`

	// 开始时间，UTC时间，例如：2024-04-03T08:00:01Z。
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间，UTC时间，例如：2024-04-03T08:00:01Z。
	EndTime *string `json:"end_time,omitempty"`

	// 对比任务的状态。取值： - RUNNING：运行中。 - WAITING_FOR_RUNNING：等待启动中。 - SUCCESSFUL：完成。 - FAILED：失败。 - CANCELLED：已取消。 - TIMEOUT_INTERRUPT：超时中断。 - FULL_DOING：全量校验中。 - INCRE_DOING：增量校验中。
	Status *string `json:"status,omitempty"`

	// 任务名称。
	JobName        *string `json:"job_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowHealthCompareJobDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHealthCompareJobDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowHealthCompareJobDetailResponse", string(data)}, " ")
}
