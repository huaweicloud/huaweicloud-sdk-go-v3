package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowQuotaResponseBodyQuotas struct {
	ProjectId *string `json:"projectId,omitempty"`

	Resources *[]ShowQuotaPropertiesBody `json:"resources,omitempty"`
}

func (o ShowQuotaResponseBodyQuotas) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQuotaResponseBodyQuotas struct{}"
	}

	return strings.Join([]string{"ShowQuotaResponseBodyQuotas", string(data)}, " ")
}
