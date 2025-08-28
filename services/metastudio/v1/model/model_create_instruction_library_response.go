package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInstructionLibraryResponse Response Object
type CreateInstructionLibraryResponse struct {

	// 指令集ID。
	InstructionLibraryId *string `json:"instruction_library_id,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateInstructionLibraryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstructionLibraryResponse struct{}"
	}

	return strings.Join([]string{"CreateInstructionLibraryResponse", string(data)}, " ")
}
