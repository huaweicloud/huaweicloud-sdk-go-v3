package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConfigValues struct {
	// 源连接参数、目的连接参数和作业任务参数，它们的配置数据结构相同，其中“inputs”里的参数不一样，详细请参见configs数据结构说明

	Configs *[]Configs `json:"configs,omitempty"`
}

func (o ConfigValues) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigValues struct{}"
	}

	return strings.Join([]string{"ConfigValues", string(data)}, " ")
}
