package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ModifyConfigReq struct {

	// **参数解释**： RocketMQ配置。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	RocketmqConfigs []RocketMqConfigReq `json:"rocketmq_configs"`
}

func (o ModifyConfigReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyConfigReq struct{}"
	}

	return strings.Join([]string{"ModifyConfigReq", string(data)}, " ")
}
