package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowTenantQuotaResponse struct {
	Absolute       *Absolute `json:"absolute,omitempty" xml:"absolute"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowTenantQuotaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTenantQuotaResponse struct{}"
	}

	return strings.Join([]string{"ShowTenantQuotaResponse", string(data)}, " ")
}
