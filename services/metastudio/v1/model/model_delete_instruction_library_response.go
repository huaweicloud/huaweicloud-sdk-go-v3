package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteInstructionLibraryResponse Response Object
type DeleteInstructionLibraryResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteInstructionLibraryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInstructionLibraryResponse struct{}"
	}

	return strings.Join([]string{"DeleteInstructionLibraryResponse", string(data)}, " ")
}
