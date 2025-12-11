package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCommonPortResponse Response Object
type ShowCommonPortResponse struct {

	// **参数解释** 端口号 **取值范围** 最小值1，最大值65535
	Port *int32 `json:"port,omitempty"`

	// **参数解释** 端口类型 **取值范围** - TCP：TCP协议 - UDP：UDP协议 - TCP6：TCP6协议 - UDP6：UDP6协议
	Type *string `json:"type,omitempty"`

	// **参数解释** 状态 **取值范围** - normal：正常 - danger：危险 - unknown：未知
	Status *string `json:"status,omitempty"`

	// **参数解释** 中文描述 **取值范围** 字符长度1-256
	Description *string `json:"description,omitempty"`

	// **参数解释** 英文描述 **取值范围** 字符长度1-256
	DescriptionEn  *string `json:"description_en,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowCommonPortResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCommonPortResponse struct{}"
	}

	return strings.Join([]string{"ShowCommonPortResponse", string(data)}, " ")
}
