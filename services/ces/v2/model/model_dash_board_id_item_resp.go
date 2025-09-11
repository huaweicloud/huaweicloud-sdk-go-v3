package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DashBoardIdItemResp struct {

	// **参数描述**： 监控看板id **取值范围** 以db开头，包含22个字母和数字，长度为24个字符
	DashboardId *string `json:"dashboard_id,omitempty"`
}

func (o DashBoardIdItemResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DashBoardIdItemResp struct{}"
	}

	return strings.Join([]string{"DashBoardIdItemResp", string(data)}, " ")
}
