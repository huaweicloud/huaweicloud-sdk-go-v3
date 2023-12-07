package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ModifyConfigReq struct {

	// RocketMQ配置。
	RocketmqConfigs *[]RocketMqConfigReq `json:"rocketmq_configs,omitempty"`
}

func (o ModifyConfigReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyConfigReq struct{}"
	}

	return strings.Join([]string{"ModifyConfigReq", string(data)}, " ")
}
