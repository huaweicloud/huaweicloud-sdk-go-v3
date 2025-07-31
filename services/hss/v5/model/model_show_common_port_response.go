package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCommonPortResponse Response Object
type ShowCommonPortResponse struct {

	// 端口号
	Port *int32 `json:"port,omitempty"`

	// 端口类型：目前包括TCP，UDP两种
	Type *string `json:"type,omitempty"`

	// 状态   - normal : 正常   - danger : 危险   - unknow : 未知
	Status *string `json:"status,omitempty"`

	// 中文描述
	Description *string `json:"description,omitempty"`

	// 英文描述
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
