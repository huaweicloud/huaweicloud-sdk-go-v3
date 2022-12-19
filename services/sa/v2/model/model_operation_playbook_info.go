package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// operation of playbook
type OperationPlaybookInfo struct {

	// 重试： RETRY  终止： TERMINATE
	Operation *string `json:"operation,omitempty"`
}

func (o OperationPlaybookInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperationPlaybookInfo struct{}"
	}

	return strings.Join([]string{"OperationPlaybookInfo", string(data)}, " ")
}
