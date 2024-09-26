package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCloudConnectionQuotasResponse Response Object
type ListCloudConnectionQuotasResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	// 配额列表。
	Quotas         []CloudConnectionQuota `json:"quotas"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListCloudConnectionQuotasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCloudConnectionQuotasResponse struct{}"
	}

	return strings.Join([]string{"ListCloudConnectionQuotasResponse", string(data)}, " ")
}
