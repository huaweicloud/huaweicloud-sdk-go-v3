package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WebHookEventCfgDto struct {

	// **参数解释：** 事件类型。 **取值范围：** 最小1个字节，最大255字节
	EventType *string `json:"event_type,omitempty"`

	// **参数解释：** 配置信息。 **取值范围：** 最小1个字节，最大255字节
	Cfgs *string `json:"cfgs,omitempty"`
}

func (o WebHookEventCfgDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WebHookEventCfgDto struct{}"
	}

	return strings.Join([]string{"WebHookEventCfgDto", string(data)}, " ")
}
