package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TablesConfig 表配置。
type TablesConfig struct {

	// 表名。
	TableName *string `json:"table_name,omitempty"`

	// 表配置值。
	TableConfig *string `json:"table_config,omitempty"`
}

func (o TablesConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TablesConfig struct{}"
	}

	return strings.Join([]string{"TablesConfig", string(data)}, " ")
}
