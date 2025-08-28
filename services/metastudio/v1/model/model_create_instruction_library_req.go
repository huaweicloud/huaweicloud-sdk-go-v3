package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInstructionLibraryReq 创建指令集请求。
type CreateInstructionLibraryReq struct {

	// 指令集名称。
	Name string `json:"name"`
}

func (o CreateInstructionLibraryReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstructionLibraryReq struct{}"
	}

	return strings.Join([]string{"CreateInstructionLibraryReq", string(data)}, " ")
}
