package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCollectorParserResponse Response Object
type ShowCollectorParserResponse struct {

	// 数值
	ChannelReferCount *int64 `json:"channel_refer_count,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 相关描述信息
	Modules *[]ShowModuleVo `json:"modules,omitempty"`

	// UUID
	ParserId *string `json:"parser_id,omitempty"`

	// 相关标题
	Title          *string `json:"title,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowCollectorParserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCollectorParserResponse struct{}"
	}

	return strings.Join([]string{"ShowCollectorParserResponse", string(data)}, " ")
}
