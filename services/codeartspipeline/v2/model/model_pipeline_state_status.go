package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PipelineStateStatus struct {

	// 阶段或任务标识
	Id string `json:"id"`

	// 阶段或任务名称
	Name string `json:"name"`

	// 类别(阶段/任务)
	Type string `json:"type"`

	// 执行开始时间
	StartTime string `json:"start_time"`

	// 执行结束时间
	EndTime string `json:"end_time"`

	// 运行耗时
	ElapsedTime string `json:"elapsed_time"`

	// 运行状态。取值和含义：waiting：等待;running：执行中;verifying：待审核；suspending：挂起;completed：完成
	Status string `json:"status"`

	// 运行结果。取值及含义：success：成功；error：失败；aborted：终止
	Outcome string `json:"outcome"`

	// 错误码
	ErrorCode string `json:"error_code"`

	// 错误信息
	ErrorMsg string `json:"error_msg"`

	// 子任务运行信息(对任务来说是空的)
	Children []PipelineStateStatus `json:"children"`

	// 任务运行记录跳转链接
	DetailUrl string `json:"detail_url"`
}

func (o PipelineStateStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipelineStateStatus struct{}"
	}

	return strings.Join([]string{"PipelineStateStatus", string(data)}, " ")
}
