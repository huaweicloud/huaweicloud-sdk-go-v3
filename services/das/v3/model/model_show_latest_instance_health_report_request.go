package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLatestInstanceHealthReportRequest Request Object
type ShowLatestInstanceHealthReportRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`
}

func (o ShowLatestInstanceHealthReportRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLatestInstanceHealthReportRequest struct{}"
	}

	return strings.Join([]string{"ShowLatestInstanceHealthReportRequest", string(data)}, " ")
}
