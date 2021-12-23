package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Configs struct {
	// 输入参数列表，列表中的每个参数为“name,value”结构，请参考inputs数据结构参数说明。在“from-config-values”数据结构中，不同的源连接类型有不同的“inputs”参数列表，请参见源端作业参数说明下的章节。在“to-cofig-values”数据结构中，不同的目的连接类型有不同的“inputs”参数列表，请参见目的端作业参数说明下面的子章节。在“driver-config-values”数据结构中，“inputs”具体参数请参见作业任务参数说明。

	Inputs []Input `json:"inputs"`
	// 配置名称：源端作业的配置名称为“fromJobConfig”。目的端作业的配置名称为“toJobConfig”,连接的配置名称固定为“linkConfig”。

	Name string `json:"name"`
}

func (o Configs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Configs struct{}"
	}

	return strings.Join([]string{"Configs", string(data)}, " ")
}
