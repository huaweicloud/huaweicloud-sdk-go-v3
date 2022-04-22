package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 回调配置
type AppCallbackUrl struct {

	// 回调通知url地址，url必须以http://或https://开头，需要支持POST调用。
	Url *string `json:"url,omitempty"`

	// 回调秘钥，主要用于鉴权。如果不设置或者为空，则回调不会增加鉴权头域字段。
	AuthKey *string `json:"auth_key,omitempty"`

	// 订阅云端录制通知消息。  取值如下：  - RECORD_NEW_FILE_START：录制模块开始创建新的录制文件。  - RECORD_FILE_COMPLETE：录制模块已经生成录制文件。
	NotifyEventSubscription *[]AppCallbackUrlNotifyEventSubscription `json:"notify_event_subscription,omitempty"`

	// 更新时间，形如“2006-01-02T15:04:05.075Z”，时区为：UTC。
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o AppCallbackUrl) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppCallbackUrl struct{}"
	}

	return strings.Join([]string{"AppCallbackUrl", string(data)}, " ")
}

type AppCallbackUrlNotifyEventSubscription struct {
	value string
}

type AppCallbackUrlNotifyEventSubscriptionEnum struct {
	RECORD_NEW_FILE_START AppCallbackUrlNotifyEventSubscription
	RECORD_FILE_COMPLETE  AppCallbackUrlNotifyEventSubscription
}

func GetAppCallbackUrlNotifyEventSubscriptionEnum() AppCallbackUrlNotifyEventSubscriptionEnum {
	return AppCallbackUrlNotifyEventSubscriptionEnum{
		RECORD_NEW_FILE_START: AppCallbackUrlNotifyEventSubscription{
			value: "RECORD_NEW_FILE_START",
		},
		RECORD_FILE_COMPLETE: AppCallbackUrlNotifyEventSubscription{
			value: "RECORD_FILE_COMPLETE",
		},
	}
}

func (c AppCallbackUrlNotifyEventSubscription) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AppCallbackUrlNotifyEventSubscription) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
