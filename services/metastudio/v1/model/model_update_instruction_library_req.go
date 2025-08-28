package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstructionLibraryReq 修改指令集请求。
type UpdateInstructionLibraryReq struct {

	// 指令集名称。
	Name *string `json:"name,omitempty"`
}

func (o UpdateInstructionLibraryReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstructionLibraryReq struct{}"
	}

	return strings.Join([]string{"UpdateInstructionLibraryReq", string(data)}, " ")
}
