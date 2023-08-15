package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListQuotaResponse Response Object
type ListQuotaResponse struct {
	Quotas         *QuotaResourcesDto `json:"quotas,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListQuotaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQuotaResponse struct{}"
	}

	return strings.Join([]string{"ListQuotaResponse", string(data)}, " ")
}
