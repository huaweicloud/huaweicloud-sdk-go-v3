package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDashboardRequest Request Object
type DeleteDashboardRequest struct {

	// 仪表盘id。
	DashboardId string `json:"dashboard_id"`

	// 企业项目id。获取方式请参见：[获取企业项目ID](aom_04_0024.xml)。 - 删除单个企业项目下实例，填写企业项目id。  - 不填时，默认删除企业项目id为0的企业项目下实例。
	EnterpriseProjectId *string `json:"Enterprise-Project-Id,omitempty"`
}

func (o DeleteDashboardRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDashboardRequest struct{}"
	}

	return strings.Join([]string{"DeleteDashboardRequest", string(data)}, " ")
}
