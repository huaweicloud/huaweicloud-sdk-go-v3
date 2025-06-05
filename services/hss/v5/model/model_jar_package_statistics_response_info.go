package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JarPackageStatisticsResponseInfo **参数解释**: 中间件包统计信息列表
type JarPackageStatisticsResponseInfo struct {

	// **参数解释**: 中间件包名称 **取值范围**: 字符长度0-256
	FileName *string `json:"file_name,omitempty"`

	// **参数解释**: 中间件包统计信息总数 **取值范围**: 取值0-300000
	Num *int32 `json:"num,omitempty"`
}

func (o JarPackageStatisticsResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JarPackageStatisticsResponseInfo struct{}"
	}

	return strings.Join([]string{"JarPackageStatisticsResponseInfo", string(data)}, " ")
}
