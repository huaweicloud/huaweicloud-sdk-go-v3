package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteTaskParams struct {

	// 测试计划id
	PlanId *string `json:"planId,omitempty"`

	// 任务id列表信息
	TaskIds *[]string `json:"taskIds,omitempty"`
}

func (o DeleteTaskParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTaskParams struct{}"
	}

	return strings.Join([]string{"DeleteTaskParams", string(data)}, " ")
}
