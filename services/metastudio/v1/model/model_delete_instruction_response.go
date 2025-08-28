package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteInstructionResponse Response Object
type DeleteInstructionResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteInstructionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInstructionResponse struct{}"
	}

	return strings.Join([]string{"DeleteInstructionResponse", string(data)}, " ")
}
