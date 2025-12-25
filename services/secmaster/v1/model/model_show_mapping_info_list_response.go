package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMappingInfoListResponse Response Object
type ShowMappingInfoListResponse struct {

	// **参数解释**: 错误码 **取值范围**: 不涉及
	Code *string `json:"code,omitempty"`

	// **参数解释**: 错误描述 **取值范围**: 不涉及
	Message *string `json:"message,omitempty"`

	// 总数
	Total *int32 `json:"total,omitempty"`

	// **参数解释：** 偏移量 **约束限制：** 0-10000 **取值范围：** 不涉及 **默认取值：** 0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**: 当前页码 **约束限制**: 不涉及
	Limit *int32 `json:"limit,omitempty"`

	// 分类映射信息集合
	Data           *[]DpeInfo `json:"data,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ShowMappingInfoListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMappingInfoListResponse struct{}"
	}

	return strings.Join([]string{"ShowMappingInfoListResponse", string(data)}, " ")
}
