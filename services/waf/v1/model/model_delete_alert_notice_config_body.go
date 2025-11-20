package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteAlertNoticeConfigBody struct {

	// **参数解释：** 告警id，用于唯一标识一条告警通知配置，id请查看”查询告警通知配置“接口 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Id *string `json:"id,omitempty"`
}

func (o DeleteAlertNoticeConfigBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAlertNoticeConfigBody struct{}"
	}

	return strings.Join([]string{"DeleteAlertNoticeConfigBody", string(data)}, " ")
}
