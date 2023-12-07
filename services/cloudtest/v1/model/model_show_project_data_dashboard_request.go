package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowProjectDataDashboardRequest Request Object
type ShowProjectDataDashboardRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	Body *TestReportInfoRequestBody `json:"body,omitempty"`
}

func (o ShowProjectDataDashboardRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProjectDataDashboardRequest struct{}"
	}

	return strings.Join([]string{"ShowProjectDataDashboardRequest", string(data)}, " ")
}
