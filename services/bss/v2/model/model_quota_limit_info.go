package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QuotaLimitInfo struct {

	// 属性key值。
	LimitKey *string `json:"limit_key,omitempty" xml:"limit_key"`

	// 属性值，具体参见表3。
	LimitValues *[]LimitValue `json:"limit_values,omitempty" xml:"limit_values"`
}

func (o QuotaLimitInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaLimitInfo struct{}"
	}

	return strings.Join([]string{"QuotaLimitInfo", string(data)}, " ")
}
