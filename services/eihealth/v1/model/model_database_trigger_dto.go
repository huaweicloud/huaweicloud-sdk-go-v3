package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DatabaseTriggerDto struct {

	// 触发器的列名，取值范围：[1,63]
	Name string `json:"name"`

	JudgeMode *TriggerJudgeMode `json:"judge_mode"`

	// 触发器的取值，取值范围：[1,128]
	Value string `json:"value"`
}

func (o DatabaseTriggerDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatabaseTriggerDto struct{}"
	}

	return strings.Join([]string{"DatabaseTriggerDto", string(data)}, " ")
}
