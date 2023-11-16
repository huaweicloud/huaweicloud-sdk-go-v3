package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBlockStatisticsRequest Request Object
type ShowBlockStatisticsRequest struct {

	// 租户id
	DomainId string `json:"domain_id"`
}

func (o ShowBlockStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBlockStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ShowBlockStatisticsRequest", string(data)}, " ")
}
