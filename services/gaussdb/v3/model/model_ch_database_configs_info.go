package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChDatabaseConfigsInfo 库配置列表。
type ChDatabaseConfigsInfo struct {

	// 库同步配置参数名。可通过“查询ClickHouse数据同步的库参数配置”接口查询。
	ParamName string `json:"param_name"`

	// 库同步配置参数值。
	Value string `json:"value"`
}

func (o ChDatabaseConfigsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChDatabaseConfigsInfo struct{}"
	}

	return strings.Join([]string{"ChDatabaseConfigsInfo", string(data)}, " ")
}
