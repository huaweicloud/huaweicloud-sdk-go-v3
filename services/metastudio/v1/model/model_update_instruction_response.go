package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstructionResponse Response Object
type UpdateInstructionResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateInstructionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstructionResponse struct{}"
	}

	return strings.Join([]string{"UpdateInstructionResponse", string(data)}, " ")
}
