package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IdcardResult struct {
	// 身份证上识别的名称。

	Name *string `json:"name,omitempty"`
	// 身份证号。

	Number *string `json:"number,omitempty"`
	// 性别。

	Sex *string `json:"sex,omitempty"`
	// 出生日期。

	Birth *string `json:"birth,omitempty"`
	// 民族。

	Ethnicity *string `json:"ethnicity,omitempty"`
	// 地址。

	Address *string `json:"address,omitempty"`
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
