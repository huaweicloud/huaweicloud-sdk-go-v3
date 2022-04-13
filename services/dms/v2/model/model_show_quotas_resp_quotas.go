package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 配额信息。
type ShowQuotasRespQuotas struct {
	// 配额列表。

	Resources *[]ShowQuotasRespQuotasResources `json:"resources,omitempty"`
}

func (o ShowQuotasRespQuotas) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQuotasRespQuotas struct{}"
	}

	return strings.Join([]string{"ShowQuotasRespQuotas", string(data)}, " ")
}
