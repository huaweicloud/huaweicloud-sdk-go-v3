package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 配额信息
type QuotasDetail struct {
	// 资源配额

	Resources []EpQuotas `json:"resources"`
}

func (o QuotasDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotasDetail struct{}"
	}

	return strings.Join([]string{"QuotasDetail", string(data)}, " ")
}
