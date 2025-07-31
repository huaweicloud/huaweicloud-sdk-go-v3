package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProcessStatisticResponseInfo 进程统计信息
type ProcessStatisticResponseInfo struct {

	// **参数解释** : 进程的可执行文件路径 **约束限制** : 不涉及 **取值范围** : 字符长度1-256位 **默认取值** : 不涉及
	Path *string `json:"path,omitempty"`

	// **参数解释** : 进程数量 **约束限制** : 不涉及 **取值范围** : 最小值0，最大值10000 **默认取值** : 不涉及
	Num *int32 `json:"num,omitempty"`
}

func (o ProcessStatisticResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProcessStatisticResponseInfo struct{}"
	}

	return strings.Join([]string{"ProcessStatisticResponseInfo", string(data)}, " ")
}
