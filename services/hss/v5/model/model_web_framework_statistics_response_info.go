package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WebFrameworkStatisticsResponseInfo **参数解释** Web框架统计信息列表
type WebFrameworkStatisticsResponseInfo struct {

	// **参数解释**: Web框架文件名称 **取值范围**: 字符长度0-256
	FileName *string `json:"file_name,omitempty"`

	// **参数解释** Web框架统计信息总数 **取值范围** 最小值0，最大值300000
	Num *int32 `json:"num,omitempty"`
}

func (o WebFrameworkStatisticsResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WebFrameworkStatisticsResponseInfo struct{}"
	}

	return strings.Join([]string{"WebFrameworkStatisticsResponseInfo", string(data)}, " ")
}
