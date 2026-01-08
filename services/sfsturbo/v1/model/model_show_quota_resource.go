package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowQuotaResource 租户配额
type ShowQuotaResource struct {

	// 租户配额
	Resources *[]QuotaResource `json:"resources,omitempty"`
}

func (o ShowQuotaResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQuotaResource struct{}"
	}

	return strings.Join([]string{"ShowQuotaResource", string(data)}, " ")
}
