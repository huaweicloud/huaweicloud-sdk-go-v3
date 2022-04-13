package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListQuotaDetailsResponse struct {
	Quotas         *ResourcesResp `json:"quotas,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListQuotaDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQuotaDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListQuotaDetailsResponse", string(data)}, " ")
}
