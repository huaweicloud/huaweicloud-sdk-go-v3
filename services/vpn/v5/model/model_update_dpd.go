package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateDpd struct {

	// 对等体存活检测空闲时间
	Interval *int32 `json:"interval,omitempty"`

	// 对等体存活检测报文重传间隔
	Timeout *int32 `json:"timeout,omitempty"`

	// 对等体存活检测报文格式
	Msg *UpdateDpdMsg `json:"msg,omitempty"`
}

func (o UpdateDpd) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDpd struct{}"
	}

	return strings.Join([]string{"UpdateDpd", string(data)}, " ")
}

type UpdateDpdMsg struct {
	value string
}

type UpdateDpdMsgEnum struct {
	SEQ_HASH_NOTIFY UpdateDpdMsg
	SEQ_NOTIFY_HASH UpdateDpdMsg
}

func GetUpdateDpdMsgEnum() UpdateDpdMsgEnum {
	return UpdateDpdMsgEnum{
		SEQ_HASH_NOTIFY: UpdateDpdMsg{
			value: "seq-hash-notify",
		},
		SEQ_NOTIFY_HASH: UpdateDpdMsg{
			value: "seq-notify-hash",
		},
	}
}

func (c UpdateDpdMsg) Value() string {
	return c.value
}

func (c UpdateDpdMsg) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateDpdMsg) UnmarshalJSON(b []byte) error {
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
