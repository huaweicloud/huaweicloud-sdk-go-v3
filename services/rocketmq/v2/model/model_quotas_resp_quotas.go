package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QuotasRespQuotas 配额信息。
type QuotasRespQuotas struct {

	// 配额列表。
	Resources *[]QuotaResourceEntity `json:"resources,omitempty"`
}

func (o QuotasRespQuotas) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotasRespQuotas struct{}"
	}

	return strings.Join([]string{"QuotasRespQuotas", string(data)}, " ")
}
