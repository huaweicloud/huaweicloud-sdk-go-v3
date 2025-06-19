package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SmartConnectValidateEntity struct {
	Task *SmartConnectTaskRespSourceConfig `json:"task,omitempty"`

	// **参数解释**： Smart Connect任务类型。 **取值范围**： - OBS_SINK：转储。 - KAFKA_REPLICATOR_SOURCE：Kafka数据复制。
	Type *string `json:"type,omitempty"`
}

func (o SmartConnectValidateEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmartConnectValidateEntity struct{}"
	}

	return strings.Join([]string{"SmartConnectValidateEntity", string(data)}, " ")
}
