package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateOneDashboardResponse Response Object
type CreateOneDashboardResponse struct {

	// **参数描述**： 监控看板id **取值范围** 以db开头，包含22个字母和数字，长度为24个字符
	DashboardId    *string `json:"dashboard_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateOneDashboardResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOneDashboardResponse struct{}"
	}

	return strings.Join([]string{"CreateOneDashboardResponse", string(data)}, " ")
}
