package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstructionLibraryResponse Response Object
type UpdateInstructionLibraryResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateInstructionLibraryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstructionLibraryResponse struct{}"
	}

	return strings.Join([]string{"UpdateInstructionLibraryResponse", string(data)}, " ")
}
