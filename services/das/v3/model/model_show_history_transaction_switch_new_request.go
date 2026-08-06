package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHistoryTransactionSwitchNewRequest Request Object
type ShowHistoryTransactionSwitchNewRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`
}

func (o ShowHistoryTransactionSwitchNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHistoryTransactionSwitchNewRequest struct{}"
	}

	return strings.Join([]string{"ShowHistoryTransactionSwitchNewRequest", string(data)}, " ")
}
