package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDashBoardRequest Request Object
type ShowDashBoardRequest struct {

	// 仪表盘id。
	DashboardId string `json:"dashboard_id"`

	// 企业项目id。获取方式请参见：[获取企业项目ID](aom_04_0024.xml)。  -  查询单个企业项目下实例，填写企业项目id。  -  查询所有企业项目下实例，填写“all_granted_eps”。
	EnterpriseProjectId *string `json:"Enterprise-Project-Id,omitempty"`

	// 仪表盘版本号。
	Version string `json:"version"`
}

func (o ShowDashBoardRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDashBoardRequest struct{}"
	}

	return strings.Join([]string{"ShowDashBoardRequest", string(data)}, " ")
}
