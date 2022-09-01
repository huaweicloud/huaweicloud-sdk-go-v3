package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 任务信息
type Task struct {

	// 任务id
	Id *string `json:"id,omitempty" xml:"id"`

	// 任务名称
	Name *string `json:"name,omitempty" xml:"name"`
}

func (o Task) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Task struct{}"
	}

	return strings.Join([]string{"Task", string(data)}, " ")
}
