package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KernelModuleStatisticsResponseInfo 内核模块统计信息列表
type KernelModuleStatisticsResponseInfo struct {

	// 内核模块名称
	Name *string `json:"name,omitempty"`

	// 内核模块统计信息总数
	Num *int32 `json:"num,omitempty"`
}

func (o KernelModuleStatisticsResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KernelModuleStatisticsResponseInfo struct{}"
	}

	return strings.Join([]string{"KernelModuleStatisticsResponseInfo", string(data)}, " ")
}
