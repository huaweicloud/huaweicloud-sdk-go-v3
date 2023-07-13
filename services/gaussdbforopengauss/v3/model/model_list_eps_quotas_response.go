package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEpsQuotasResponse Response Object
type ListEpsQuotasResponse struct {

	// 企业配额列表。
	EpsQuotas *[]ListQuotaResult `json:"eps_quotas,omitempty"`

	// 配额组数量。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListEpsQuotasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEpsQuotasResponse struct{}"
	}

	return strings.Join([]string{"ListEpsQuotasResponse", string(data)}, " ")
}
