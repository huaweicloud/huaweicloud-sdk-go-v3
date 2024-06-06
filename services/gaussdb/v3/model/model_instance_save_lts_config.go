package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InstanceSaveLtsConfig struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 日志类型。
	LogType string `json:"log_type"`

	// 日志组ID。
	LtsGroupId string `json:"lts_group_id"`

	// 日志流ID。
	LtsStreamId string `json:"lts_stream_id"`
}

func (o InstanceSaveLtsConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceSaveLtsConfig struct{}"
	}

	return strings.Join([]string{"InstanceSaveLtsConfig", string(data)}, " ")
}
