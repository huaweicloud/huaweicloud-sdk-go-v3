package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDashboardInfosRequest Request Object
type ListDashboardInfosRequest struct {

	// 企业项目Id
	EnterpriseId *string `json:"enterprise_id,omitempty"`

	// 指定企业项目下监控面板是否收藏，true:收藏，false:未收藏，填此参数时，enterprise_id必填
	IsFavorite *bool `json:"is_favorite,omitempty"`

	// 监控面板名称
	DashboardName *string `json:"dashboard_name,omitempty"`

	// 监控面板id
	DashboardId *string `json:"dashboard_id,omitempty"`
}

func (o ListDashboardInfosRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDashboardInfosRequest struct{}"
	}

	return strings.Join([]string{"ListDashboardInfosRequest", string(data)}, " ")
}
