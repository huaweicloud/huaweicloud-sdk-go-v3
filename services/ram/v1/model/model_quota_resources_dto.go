package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QuotaResourcesDto 配额信息列表。
type QuotaResourcesDto struct {

	// 配额信息。
	Resources []Quotas `json:"resources"`
}

func (o QuotaResourcesDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaResourcesDto struct{}"
	}

	return strings.Join([]string{"QuotaResourcesDto", string(data)}, " ")
}
