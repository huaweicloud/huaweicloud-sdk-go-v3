package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowJudgementDetailResponse struct {

	// 任务回调状态:callback_success(回调成功)、callback_error(回调出错)
	TaskStatus *string `json:"task_status,omitempty" xml:"task_status"`

	// 任务运行状态:success(成功)、compile_error(编译错误)、internal_error(系统内部错误)、run_timeout(运行超时)、judging(代码运行中)、runtime_error(运行出错)、output_match_error(不符合输出规范)
	Status *string `json:"status,omitempty" xml:"status"`

	// 判题任务ID
	JudgementId *string `json:"judgement_id,omitempty" xml:"judgement_id"`

	// 任务开始时间
	StartTime *string `json:"start_time,omitempty" xml:"start_time"`

	// 任务结束时间
	EndTime *string `json:"end_time,omitempty" xml:"end_time"`

	Result         *JudgementResult `json:"result,omitempty" xml:"result"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowJudgementDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJudgementDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowJudgementDetailResponse", string(data)}, " ")
}
