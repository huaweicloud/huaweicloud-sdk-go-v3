package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AutoRecordModeReq 配置app自动录制模式请求
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

func (c AutoRecordModeReqMode) Value() string {
	return c.value
}

func (c AutoRecordModeReqMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AutoRecordModeReqMode) UnmarshalJSON(b []byte) error {
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
