package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AlertNoticeConfigList struct {

	// **参数解释：** 告警通知配置列表，包含多条告警通知配置信息 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	AlertNoticeConfigs []AlertNoticeConfig `json:"alert_notice_configs"`
}

func (o AlertNoticeConfigList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlertNoticeConfigList struct{}"
	}

	return strings.Join([]string{"AlertNoticeConfigList", string(data)}, " ")
}
