package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListQuotaDetailsResponse Response Object
type ListQuotaDetailsResponse struct {

	// 请求ID。  注：自动生成 。
	RequestId *string `json:"request_id,omitempty"`

	// 资源配额信息列表。
	Quotas         *[]QuotaInfo `json:"quotas,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListQuotaDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQuotaDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListQuotaDetailsResponse", string(data)}, " ")
}
