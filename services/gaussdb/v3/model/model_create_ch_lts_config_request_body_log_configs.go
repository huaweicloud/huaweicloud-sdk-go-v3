package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateChLtsConfigRequestBodyLogConfigs struct {

	// ClickHouse实例ID，严格匹配UUID规则。
	InstanceId string `json:"instance_id"`

	// 查询日志类型。当前仅支持slow_log。
	LogType string `json:"log_type"`

	// LTS日志组id。
	LtsGroupId string `json:"lts_group_id"`

	// LTS日志流id。
	LtsStreamId string `json:"lts_stream_id"`
}

func (o CreateChLtsConfigRequestBodyLogConfigs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateChLtsConfigRequestBodyLogConfigs struct{}"
	}

	return strings.Join([]string{"CreateChLtsConfigRequestBodyLogConfigs", string(data)}, " ")
}
