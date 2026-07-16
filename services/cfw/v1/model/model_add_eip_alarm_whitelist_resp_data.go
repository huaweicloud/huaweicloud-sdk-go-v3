package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddEipAlarmWhitelistRespData struct {

	// **参数解释**： 防火墙实例id **约束限制**： 不涉及
	Id *string `json:"id,omitempty"`
}

func (o AddEipAlarmWhitelistRespData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddEipAlarmWhitelistRespData struct{}"
	}

	return strings.Join([]string{"AddEipAlarmWhitelistRespData", string(data)}, " ")
}
