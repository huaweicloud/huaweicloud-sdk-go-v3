package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstructionRequest Request Object
type ListInstructionRequest struct {

	// 使用AK/SK方式认证时必选，携带的鉴权信息。
	Authorization *string `json:"Authorization,omitempty"`

	// 使用AK/SK方式认证时必选，请求的发生时间。
	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	// 使用AK/SK方式认证时必选，携带项目ID信息。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 第三方用户ID。不允许输入中文。
	XAppUserId *string `json:"X-App-UserId,omitempty"`

	// 偏移量，表示从此偏移量开始查询。
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty"`

	// 指令集ID。
	InstructionLibraryId string `json:"instruction_library_id"`

	// 指令名称
	Name *string `json:"name,omitempty"`
}

func (o ListInstructionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstructionRequest struct{}"
	}

	return strings.Join([]string{"ListInstructionRequest", string(data)}, " ")
}
