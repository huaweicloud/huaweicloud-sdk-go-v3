package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KernelModuleStatisticsResponseInfo **参数解释** 内核模块统计信息列表
type KernelModuleStatisticsResponseInfo struct {

	// **参数解释**: 内核模块名称 **取值范围**: 字符长度0-256
	Name *string `json:"name,omitempty"`

	// **参数解释** 内核模块统计信息总数 **取值范围** 最小值0，最大值300000
	Num *int32 `json:"num,omitempty"`
}

func (o KernelModuleStatisticsResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KernelModuleStatisticsResponseInfo struct{}"
	}

	return strings.Join([]string{"KernelModuleStatisticsResponseInfo", string(data)}, " ")
}
