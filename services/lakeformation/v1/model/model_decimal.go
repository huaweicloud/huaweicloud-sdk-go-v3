package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Decimal struct {

	// 整数部分
	Scale *int32 `json:"scale,omitempty"`

	// 小数部分
	Unscaled string `json:"unscaled,omitempty"`
}

func (o Decimal) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Decimal struct{}"
	}

	return strings.Join([]string{"Decimal", string(data)}, " ")
}
