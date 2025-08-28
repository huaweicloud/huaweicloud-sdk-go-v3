package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInstructionResponse Response Object
type CreateInstructionResponse struct {

	// 指令ID。
	InstructionId *string `json:"instruction_id,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateInstructionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstructionResponse struct{}"
	}

	return strings.Join([]string{"CreateInstructionResponse", string(data)}, " ")
}
