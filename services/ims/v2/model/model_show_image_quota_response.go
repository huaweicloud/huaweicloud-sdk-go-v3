package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowImageQuotaResponse struct {
	Quotas         *Quota `json:"quotas,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowImageQuotaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowImageQuotaResponse struct{}"
	}

	return strings.Join([]string{"ShowImageQuotaResponse", string(data)}, " ")
}
