package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEpsQuotasResponse Response Object
type ListEpsQuotasResponse struct {

	// 总记录数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 企业项目配额信息列表。
	Quotas         *[]NoSqlQueryEpsQuotaInfo `json:"quotas,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListEpsQuotasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEpsQuotasResponse struct{}"
	}

	return strings.Join([]string{"ListEpsQuotasResponse", string(data)}, " ")
}
