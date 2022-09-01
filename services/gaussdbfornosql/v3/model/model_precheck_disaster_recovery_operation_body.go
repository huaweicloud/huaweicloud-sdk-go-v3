package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PrecheckDisasterRecoveryOperationBody struct {

	// 指定预校验的具体容灾操作。
	Operation string `json:"operation" xml:"operation"`

	DisasterRecoveryInstance *PrecheckDisasterRecoveryInstance `json:"disaster_recovery_instance,omitempty" xml:"disaster_recovery_instance"`
}

func (o PrecheckDisasterRecoveryOperationBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrecheckDisasterRecoveryOperationBody struct{}"
	}

	return strings.Join([]string{"PrecheckDisasterRecoveryOperationBody", string(data)}, " ")
}
