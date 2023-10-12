package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowQuotasResponseBodyQuotas 配额使用详情
type ShowQuotasResponseBodyQuotas struct {
	Resources *[]Info `json:"resources,omitempty"`
}

func (o ShowQuotasResponseBodyQuotas) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQuotasResponseBodyQuotas struct{}"
	}

	return strings.Join([]string{"ShowQuotasResponseBodyQuotas", string(data)}, " ")
}
