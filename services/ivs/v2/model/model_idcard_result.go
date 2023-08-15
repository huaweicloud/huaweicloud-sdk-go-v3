package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IdcardResult struct {

	// 身份证上识别的名称。
	Name string `json:"name"`

	// 身份证号。
	Number string `json:"number"`

	// 性别。
	Sex string `json:"sex"`

	// 出生日期。
	Birth string `json:"birth"`

	// 民族。
	Ethnicity string `json:"ethnicity"`

	// 地址。
	Address string `json:"address"`

	// 发证机关。
	Issue *string `json:"issue,omitempty"`

	// 有效起始日期。
	ValidFrom *string `json:"valid_from,omitempty"`

	// 有效结束日期。
	ValidTo *string `json:"valid_to,omitempty"`
}

func (o IdcardResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IdcardResult struct{}"
	}

	return strings.Join([]string{"IdcardResult", string(data)}, " ")
}
