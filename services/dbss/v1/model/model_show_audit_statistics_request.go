package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAuditStatisticsRequest Request Object
type ShowAuditStatisticsRequest struct {
}

func (o ShowAuditStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAuditStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ShowAuditStatisticsRequest", string(data)}, " ")
}
