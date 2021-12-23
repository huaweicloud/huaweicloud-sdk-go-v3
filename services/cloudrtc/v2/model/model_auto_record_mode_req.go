package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 配置app自动录制模式请求
type AutoRecordModeReq struct {
	// 录制模式。 - AUTO_RECORD_OFF：关闭自动录制。 - AUTO_INDIVIDUAL_RECORD：开启单流自动录制，此时record_rule_id必须非空。

	Mode AutoRecordModeReqMode `json:"mode"`
	// 录制规则id。

	RecordRuleId *string `json:"record_rule_id,omitempty"`
}

func (o AutoRecordModeReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AutoRecordModeReq struct{}"
	}

	return strings.Join([]string{"AutoRecordModeReq", string(data)}, " ")
}

type AutoRecordModeReqMode struct {
	value string
}

type AutoRecordModeReqModeEnum struct {
	AUTO_RECORD_OFF        AutoRecordModeReqMode
	AUTO_INDIVIDUAL_RECORD AutoRecordModeReqMode
}

func GetAutoRecordModeReqModeEnum() AutoRecordModeReqModeEnum {
	return AutoRecordModeReqModeEnum{
		AUTO_RECORD_OFF: AutoRecordModeReqMode{
			value: "AUTO_RECORD_OFF",
		},
		AUTO_INDIVIDUAL_RECORD: AutoRecordModeReqMode{
			value: "AUTO_INDIVIDUAL_RECORD",
		},
	}
}

func (c AutoRecordModeReqMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AutoRecordModeReqMode) UnmarshalJSON(b []byte) error {
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
