package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ManagedOperationTypeHolder struct {
	ManagedOperation *ManagedOperation `json:"managed_operation,omitempty"`
}

func (o ManagedOperationTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ManagedOperationTypeHolder struct{}"
	}

	return strings.Join([]string{"ManagedOperationTypeHolder", string(data)}, " ")
}
