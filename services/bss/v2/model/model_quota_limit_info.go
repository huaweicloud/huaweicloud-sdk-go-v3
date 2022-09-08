package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QuotaLimitInfo struct {

	// 属性key值。
	LimitKey *string `json:"limit_key,omitempty"`

	// 属性值，具体参见表3。
	LimitValues *[]LimitValue `json:"limit_values,omitempty"`
}

func (o QuotaLimitInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaLimitInfo struct{}"
	}

	return strings.Join([]string{"QuotaLimitInfo", string(data)}, " ")
}
