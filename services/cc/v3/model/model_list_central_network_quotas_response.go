package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCentralNetworkQuotasResponse Response Object
type ListCentralNetworkQuotasResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	// 配额列表
	Quotas         []CentralNetworkQuota `json:"quotas"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListCentralNetworkQuotasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCentralNetworkQuotasResponse struct{}"
	}

	return strings.Join([]string{"ListCentralNetworkQuotasResponse", string(data)}, " ")
}
