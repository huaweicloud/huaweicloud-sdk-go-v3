package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DataBaseInfo 数据库校验信息。
type DataBaseInfo struct {

	// 数据库名称。
	DatabaseName *string `json:"database_name,omitempty"`

	// 数据库配置检查结果。
	DbConfigCheckResults *[]DbConfigCheckResult `json:"db_config_check_results,omitempty"`
}

func (o DataBaseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataBaseInfo struct{}"
	}

	return strings.Join([]string{"DataBaseInfo", string(data)}, " ")
}
