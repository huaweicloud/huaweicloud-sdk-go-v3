package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetLongHistoryTransactionSwitchNewRequest Request Object
type SetLongHistoryTransactionSwitchNewRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *SetLongHistoryTransactionSwitchNewRequestBody `json:"body,omitempty"`
}

func (o SetLongHistoryTransactionSwitchNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetLongHistoryTransactionSwitchNewRequest struct{}"
	}

	return strings.Join([]string{"SetLongHistoryTransactionSwitchNewRequest", string(data)}, " ")
}
