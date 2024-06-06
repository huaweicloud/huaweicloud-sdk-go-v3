package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InstanceDeleteLtsConfig struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 日志类型
	LogType string `json:"log_type"`
}

func (o InstanceDeleteLtsConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceDeleteLtsConfig struct{}"
	}

	return strings.Join([]string{"InstanceDeleteLtsConfig", string(data)}, " ")
}
