package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteChLtsConfigRequestBodyLogConfigs struct {

	// ClickHouse实例ID，严格匹配UUID规则。
	InstanceId string `json:"instance_id"`

	// 日志类型。当前仅支持slow_log。
	LogType string `json:"log_type"`
}

func (o DeleteChLtsConfigRequestBodyLogConfigs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteChLtsConfigRequestBodyLogConfigs struct{}"
	}

	return strings.Join([]string{"DeleteChLtsConfigRequestBodyLogConfigs", string(data)}, " ")
}
