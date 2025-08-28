package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstructionLibraryResponse Response Object
type ShowInstructionLibraryResponse struct {

	// 指令集ID。
	InstructionLibraryId *string `json:"instruction_library_id,omitempty"`

	// 指令集名称。
	Name *string `json:"name,omitempty"`

	// 创建时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	UpdateTime *string `json:"update_time,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowInstructionLibraryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstructionLibraryResponse struct{}"
	}

	return strings.Join([]string{"ShowInstructionLibraryResponse", string(data)}, " ")
}
