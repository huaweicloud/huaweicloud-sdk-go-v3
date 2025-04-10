package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRaspProtectStatisticsRequest Request Object
type ShowRaspProtectStatisticsRequest struct {

	// 企业项目ID
	EnterpriseProjectId string `json:"enterprise_project_id"`
}

func (o ShowRaspProtectStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRaspProtectStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ShowRaspProtectStatisticsRequest", string(data)}, " ")
}
