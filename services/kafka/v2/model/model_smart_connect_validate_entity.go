package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SmartConnectValidateEntity struct {
	Task *SmartConnectTaskRespSourceConfig `json:"task,omitempty"`

	// SmartConnect任务类型。
	Type *string `json:"type,omitempty"`
}

func (o SmartConnectValidateEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmartConnectValidateEntity struct{}"
	}

	return strings.Join([]string{"SmartConnectValidateEntity", string(data)}, " ")
}
