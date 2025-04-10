package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type DisasterRecoverStartXlogKeepRequestBody struct {

	// 日志保留空间占可使用剩余磁盘容量的比例, 可设置范围为1-100。 properties:
	XlogKeepRatio *int32 `json:"xlog_keep_ratio,omitempty"`

	// 容灾类型。
	DisasterType DisasterRecoverStartXlogKeepRequestBodyDisasterType `json:"disaster_type"`
}

func (o DisasterRecoverStartXlogKeepRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisasterRecoverStartXlogKeepRequestBody struct{}"
	}

	return strings.Join([]string{"DisasterRecoverStartXlogKeepRequestBody", string(data)}, " ")
}

type DisasterRecoverStartXlogKeepRequestBodyDisasterType struct {
	value string
}

type DisasterRecoverStartXlogKeepRequestBodyDisasterTypeEnum struct {
	STREAM DisasterRecoverStartXlogKeepRequestBodyDisasterType
}

func GetDisasterRecoverStartXlogKeepRequestBodyDisasterTypeEnum() DisasterRecoverStartXlogKeepRequestBodyDisasterTypeEnum {
	return DisasterRecoverStartXlogKeepRequestBodyDisasterTypeEnum{
		STREAM: DisasterRecoverStartXlogKeepRequestBodyDisasterType{
			value: "stream",
		},
	}
}

func (c DisasterRecoverStartXlogKeepRequestBodyDisasterType) Value() string {
	return c.value
}

func (c DisasterRecoverStartXlogKeepRequestBodyDisasterType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DisasterRecoverStartXlogKeepRequestBodyDisasterType) UnmarshalJSON(b []byte) error {
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
