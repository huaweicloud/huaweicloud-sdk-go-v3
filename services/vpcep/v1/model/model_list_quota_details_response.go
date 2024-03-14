package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListQuotaDetailsResponse Response Object
type ListQuotaDetailsResponse struct {
	Quotas         *ResourcesResponseBody `json:"quotas,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListQuotaDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQuotaDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListQuotaDetailsResponse", string(data)}, " ")
}
