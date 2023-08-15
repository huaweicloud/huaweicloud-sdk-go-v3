package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NextflowTaskResourceRequested struct {

	// 容器名称
	Container *string `json:"container,omitempty"`

	// 执行队列，使用','分隔多个值
	Queue *string `json:"queue,omitempty"`

	// 指定task执行需要的cpu数量
	Cpus *int32 `json:"cpus,omitempty"`

	// 指定task执行需要的内存大小
	Memory *string `json:"memory,omitempty"`

	// 指定task执行需要的磁盘大小
	Disk *string `json:"disk,omitempty"`

	// 指定task执行需要的时间
	Time *string `json:"time,omitempty"`
}

func (o NextflowTaskResourceRequested) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NextflowTaskResourceRequested struct{}"
	}

	return strings.Join([]string{"NextflowTaskResourceRequested", string(data)}, " ")
}
