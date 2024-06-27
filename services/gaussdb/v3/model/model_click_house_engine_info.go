package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClickHouseEngineInfo 引擎信息。
type ClickHouseEngineInfo struct {

	// 数据库引擎类型，现在只支持click-house。
	Type string `json:"type"`

	// 数据库版本，可通过“HTAP查询引擎信息”获取。 一位数的大版本号。
	Version string `json:"version"`
}

func (o ClickHouseEngineInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClickHouseEngineInfo struct{}"
	}

	return strings.Join([]string{"ClickHouseEngineInfo", string(data)}, " ")
}
