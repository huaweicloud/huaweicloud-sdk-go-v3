package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetHistoryTransactionSwitchNewRequest Request Object
type SetHistoryTransactionSwitchNewRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *SetHistoryTransactionSwitchNewRequestBody `json:"body,omitempty"`
}

func (o SetHistoryTransactionSwitchNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetHistoryTransactionSwitchNewRequest struct{}"
	}

	return strings.Join([]string{"SetHistoryTransactionSwitchNewRequest", string(data)}, " ")
}
