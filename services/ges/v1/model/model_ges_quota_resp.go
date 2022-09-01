package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GesQuotaResp struct {

	// GES资源配额列表。
	Resources *[]Quota `json:"resources,omitempty" xml:"resources"`
}

func (o GesQuotaResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GesQuotaResp struct{}"
	}

	return strings.Join([]string{"GesQuotaResp", string(data)}, " ")
}
