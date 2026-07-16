package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetHyperinstanceOperationResponse Response Object
type GetHyperinstanceOperationResponse struct {

	// **参数解释**：操作ID。 **取值范围**：长度为[8,36]个字符。
	OperationId *string `json:"operation_id,omitempty"`

	// **参数解释**：操作状态。 **取值范围**：长度为[8,36]个字符。
	OperationStatus *string `json:"operation_status,omitempty"`

	// **参数解释**：操作类型。 **取值范围**：长度为[8,36]个字符。
	OperationType *string `json:"operation_type,omitempty"`

	OperationError *ServerOperationError `json:"operation_error,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o GetHyperinstanceOperationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetHyperinstanceOperationResponse struct{}"
	}

	return strings.Join([]string{"GetHyperinstanceOperationResponse", string(data)}, " ")
}
