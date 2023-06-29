package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QuotasResourcesDto 组织配额的响应体。
type QuotasResourcesDto struct {

	// 配额信息。
	Resources []QuotaDto `json:"resources"`
}

func (o QuotasResourcesDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotasResourcesDto struct{}"
	}

	return strings.Join([]string{"QuotasResourcesDto", string(data)}, " ")
}
