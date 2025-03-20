package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Operation 操作。
type Operation struct {

	// OpenAPI的操作标识符。
	OperationId string `json:"operation_id"`

	// 三段式的授权项名称，例如\"iam:policies:createV5\"。
	OperationAction string `json:"operation_action"`

	// 该操作可能需要的其他授权项授权。
	DependentActions *[]string `json:"dependent_actions,omitempty"`
}

func (o Operation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Operation struct{}"
	}

	return strings.Join([]string{"Operation", string(data)}, " ")
}
