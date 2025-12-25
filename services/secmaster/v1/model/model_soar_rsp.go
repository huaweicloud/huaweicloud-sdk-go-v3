package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SoarRsp 请求的返回对象
type SoarRsp struct {

	// **参数解释**: 错误码 **取值范围**: 不涉及
	Code *string `json:"code,omitempty"`

	// **参数解释**: 错误描述 **取值范围**: 不涉及
	Message *string `json:"message,omitempty"`

	// **参数解释**: 是否成功 **取值范围**: - true  成功 - false 失败
	Success *bool `json:"success,omitempty"`

	// **参数解释**: 请求的ID **约束限制**: 不涉及
	RequestId *string `json:"request_id,omitempty"`

	// **参数解释**: 数据总条数 **约束限制**: 不涉及
	Total *int32 `json:"total,omitempty"`

	// **参数解释**: 当前页大小 **约束限制**: 不涉及
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**: 当前页码 **约束限制**: 不涉及
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**: 当前页大小 **约束限制**: 不涉及
	Size *int32 `json:"size,omitempty"`

	// **参数解释**: 分页的页数 **约束限制**: 不涉及
	Page *int32 `json:"page,omitempty"`
}

func (o SoarRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SoarRsp struct{}"
	}

	return strings.Join([]string{"SoarRsp", string(data)}, " ")
}
