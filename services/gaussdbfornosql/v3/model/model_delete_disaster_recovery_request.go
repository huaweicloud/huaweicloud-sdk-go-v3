package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDisasterRecoveryRequest Request Object
type DeleteDisasterRecoveryRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`
}

func (o DeleteDisasterRecoveryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDisasterRecoveryRequest struct{}"
	}

	return strings.Join([]string{"DeleteDisasterRecoveryRequest", string(data)}, " ")
}
