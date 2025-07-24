package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEventResponse Response Object
type ShowEventResponse struct {

	// 事件码
	Code *string `json:"code,omitempty"`

	// 状态
	Status *string `json:"status,omitempty"`

	// 事件中文名称
	NameCn *string `json:"name_cn,omitempty"`

	// 事件英文名称
	NameEn *string `json:"name_en,omitempty"`

	// 事件中文影响
	EffectCn *string `json:"effect_cn,omitempty"`

	// 事件英文影响
	EffectEn *string `json:"effect_en,omitempty"`

	// 事件中文处理建议
	SuggestionCn *string `json:"suggestion_cn,omitempty"`

	// 事件英文处理建议
	SuggestionEn *string `json:"suggestion_en,omitempty"`

	// 事件中文原因
	ReasonCn *string `json:"reason_cn,omitempty"`

	// 事件英文原因
	ReasonEn       *string `json:"reason_en,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowEventResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEventResponse struct{}"
	}

	return strings.Join([]string{"ShowEventResponse", string(data)}, " ")
}
