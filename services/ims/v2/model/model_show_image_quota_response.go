package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowImageQuotaResponse struct {
	Quotas         *Quota `json:"quotas,omitempty" xml:"quotas"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowImageQuotaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowImageQuotaResponse struct{}"
	}

	return strings.Join([]string{"ShowImageQuotaResponse", string(data)}, " ")
}
