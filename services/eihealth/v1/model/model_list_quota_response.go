package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListQuotaResponse Response Object
type ListQuotaResponse struct {

	// 配额信息列表
	Quotas *[]QuotaRsp `json:"quotas,omitempty"`

	// 配额列表个数
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListQuotaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQuotaResponse struct{}"
	}

	return strings.Join([]string{"ListQuotaResponse", string(data)}, " ")
}
