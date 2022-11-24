package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AttributeReq struct {

	// SIM卡标识
	SimCardId int64 `json:"sim_card_id"`

	// 自定义属性一
	CustomerAttribute1 *string `json:"customer_attribute1,omitempty"`

	// 自定义属性二
	CustomerAttribute2 *string `json:"customer_attribute2,omitempty"`

	// 自定义属性三
	CustomerAttribute3 *string `json:"customer_attribute3,omitempty"`

	// 自定义属性四
	CustomerAttribute4 *string `json:"customer_attribute4,omitempty"`

	// 自定义属性五
	CustomerAttribute5 *string `json:"customer_attribute5,omitempty"`

	// 自定义属性六
	CustomerAttribute6 *string `json:"customer_attribute6,omitempty"`
}

func (o AttributeReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttributeReq struct{}"
	}

	return strings.Join([]string{"AttributeReq", string(data)}, " ")
}
