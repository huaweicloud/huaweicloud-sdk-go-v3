package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PreviewStatusReason 提供有关访问预览当前状态的更多详细信息。
type PreviewStatusReason struct {

	// 访问预览当前状态的原因。
	Code PreviewStatusReasonCode `json:"code"`
}

func (o PreviewStatusReason) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PreviewStatusReason struct{}"
	}

	return strings.Join([]string{"PreviewStatusReason", string(data)}, " ")
}

type PreviewStatusReasonCode struct {
	value string
}

type PreviewStatusReasonCodeEnum struct {
	INTERNAL_ERROR        PreviewStatusReasonCode
	INVALID_CONFIGURATION PreviewStatusReasonCode
}

func GetPreviewStatusReasonCodeEnum() PreviewStatusReasonCodeEnum {
	return PreviewStatusReasonCodeEnum{
		INTERNAL_ERROR: PreviewStatusReasonCode{
			value: "internal_error",
		},
		INVALID_CONFIGURATION: PreviewStatusReasonCode{
			value: "invalid_configuration",
		},
	}
}

func (c PreviewStatusReasonCode) Value() string {
	return c.value
}

func (c PreviewStatusReasonCode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PreviewStatusReasonCode) UnmarshalJSON(b []byte) error {
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
