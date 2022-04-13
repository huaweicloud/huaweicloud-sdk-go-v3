package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 批量停止边缘实例对象
type BatchStop struct {
	// 待停止的边缘实例列表。

	Servers []BaseId `json:"servers"`
	// 关机类型，默认为SOFT。  取值范围： - SOFT：普通关机。 - HARD：强制关机。

	Type *string `json:"type,omitempty"`
}

func (o BatchStop) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStop struct{}"
	}

	return strings.Join([]string{"BatchStop", string(data)}, " ")
}
