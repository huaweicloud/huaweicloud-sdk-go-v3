package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstructionLibraryResponse Response Object
type ListInstructionLibraryResponse struct {

	// 与第一条数据的偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 页面大小
	Limit *int32 `json:"limit,omitempty"`

	// 总数量
	Count *int32 `json:"count,omitempty"`

	// 指令集信息
	Data *[]InstructionLibraryInfo `json:"data,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListInstructionLibraryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstructionLibraryResponse struct{}"
	}

	return strings.Join([]string{"ListInstructionLibraryResponse", string(data)}, " ")
}
