package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUnblockQuotaStatisticsRequest Request Object
type ListUnblockQuotaStatisticsRequest struct {

	// 租户id
	DomainId string `json:"domain_id"`
}

func (o ListUnblockQuotaStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUnblockQuotaStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ListUnblockQuotaStatisticsRequest", string(data)}, " ")
}
