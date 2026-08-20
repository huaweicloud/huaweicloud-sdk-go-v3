package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceQuotasResponse Response Object
type ShowInstanceQuotasResponse struct {
	Quotas         *InstanceQuotaResult `json:"quotas,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ShowInstanceQuotasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceQuotasResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceQuotasResponse", string(data)}, " ")
}
