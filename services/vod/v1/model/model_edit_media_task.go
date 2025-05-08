package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EditMediaTask struct {

	// 剪辑任务的输入文件信息
	Inputs *[]EditMediaTaskInput `json:"inputs,omitempty"`

	Output *TaskOutPut `json:"output,omitempty"`
}

func (o EditMediaTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EditMediaTask struct{}"
	}

	return strings.Join([]string{"EditMediaTask", string(data)}, " ")
}
