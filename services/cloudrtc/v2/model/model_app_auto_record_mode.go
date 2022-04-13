package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// app自动录制模式
type AppAutoRecordMode struct {
	// 录制模式。 - AUTO_RECORD_OFF：关闭自动录制。 - AUTO_INDIVIDUAL_RECORD：开启单流自动录制，此时record_rule_id必须非空。

	Mode AppAutoRecordModeMode `json:"mode"`
	// 录制规则id。

	RecordRuleId *string `json:"record_rule_id,omitempty"`
	// 更新时间，形如“2006-01-02T15:04:05.075Z”，时区为：UTC。

	UpdateTime *string `json:"update_time,omitempty"`
}

func (o AppAutoRecordMode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppAutoRecordMode struct{}"
	}

	return strings.Join([]string{"AppAutoRecordMode", string(data)}, " ")
}

type AppAutoRecordModeMode struct {
	value string
}

type AppAutoRecordModeModeEnum struct {
	AUTO_RECORD_OFF        AppAutoRecordModeMode
	AUTO_INDIVIDUAL_RECORD AppAutoRecordModeMode
}

func GetAppAutoRecordModeModeEnum() AppAutoRecordModeModeEnum {
	return AppAutoRecordModeModeEnum{
		AUTO_RECORD_OFF: AppAutoRecordModeMode{
			value: "AUTO_RECORD_OFF",
		},
		AUTO_INDIVIDUAL_RECORD: AppAutoRecordModeMode{
			value: "AUTO_INDIVIDUAL_RECORD",
		},
	}
}

func (c AppAutoRecordModeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AppAutoRecordModeMode) UnmarshalJSON(b []byte) error {
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
