package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// LiveNotifyConfigReq 直播通知配置。通过短信，邮件，站内信通知用户直播中断。用户可在华为云消息中心配置接受者信息
type LiveNotifyConfigReq struct {

	// 确认操作。 * add: 增加。 * del: 删除。 * replace：替换。
	Action *LiveNotifyConfigReqAction `json:"action,omitempty"`

	// **参数解释**： 启用通知事件列表。 **约束限制**： 不涉及 **默认取值**： 无。
	NotifyEvents *[]NotifyEventEnum `json:"notify_events,omitempty"`
}

func (o LiveNotifyConfigReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LiveNotifyConfigReq struct{}"
	}

	return strings.Join([]string{"LiveNotifyConfigReq", string(data)}, " ")
}

type LiveNotifyConfigReqAction struct {
	value string
}

type LiveNotifyConfigReqActionEnum struct {
	ADD     LiveNotifyConfigReqAction
	DEL     LiveNotifyConfigReqAction
	REPLACE LiveNotifyConfigReqAction
}

func GetLiveNotifyConfigReqActionEnum() LiveNotifyConfigReqActionEnum {
	return LiveNotifyConfigReqActionEnum{
		ADD: LiveNotifyConfigReqAction{
			value: "add",
		},
		DEL: LiveNotifyConfigReqAction{
			value: "del",
		},
		REPLACE: LiveNotifyConfigReqAction{
			value: "replace",
		},
	}
}

func (c LiveNotifyConfigReqAction) Value() string {
	return c.value
}

func (c LiveNotifyConfigReqAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LiveNotifyConfigReqAction) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
