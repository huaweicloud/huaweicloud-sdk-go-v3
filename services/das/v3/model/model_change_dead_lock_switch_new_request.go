package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeDeadLockSwitchNewRequest Request Object
type ChangeDeadLockSwitchNewRequest struct {
	Body *ChangeDeadLockSwitchNewRequestBody `json:"body,omitempty"`
}

func (o ChangeDeadLockSwitchNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeDeadLockSwitchNewRequest struct{}"
	}

	return strings.Join([]string{"ChangeDeadLockSwitchNewRequest", string(data)}, " ")
}
