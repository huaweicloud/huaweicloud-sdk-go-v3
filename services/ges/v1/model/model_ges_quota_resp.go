package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GesQuotaResp struct {
	// GES资源配额列表。

	Resources *[]Quota `json:"resources,omitempty"`
}

func (o GesQuotaResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GesQuotaResp struct{}"
	}

	return strings.Join([]string{"GesQuotaResp", string(data)}, " ")
}
