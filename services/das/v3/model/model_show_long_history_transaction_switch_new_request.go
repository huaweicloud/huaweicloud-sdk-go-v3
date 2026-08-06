package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLongHistoryTransactionSwitchNewRequest Request Object
type ShowLongHistoryTransactionSwitchNewRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`
}

func (o ShowLongHistoryTransactionSwitchNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLongHistoryTransactionSwitchNewRequest struct{}"
	}

	return strings.Join([]string{"ShowLongHistoryTransactionSwitchNewRequest", string(data)}, " ")
}
