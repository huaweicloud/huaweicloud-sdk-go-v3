package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WebFrameworkStatisticsResponseInfo Web框架统计信息列表
type WebFrameworkStatisticsResponseInfo struct {

	// Web框架文件名称
	FileName *string `json:"file_name,omitempty"`

	// Web框架统计信息总数
	Num *int32 `json:"num,omitempty"`
}

func (o WebFrameworkStatisticsResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WebFrameworkStatisticsResponseInfo struct{}"
	}

	return strings.Join([]string{"WebFrameworkStatisticsResponseInfo", string(data)}, " ")
}
