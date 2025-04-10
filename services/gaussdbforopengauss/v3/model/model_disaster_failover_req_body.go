package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DisasterFailoverReqBody 容灾灾备升主请求参数。
type DisasterFailoverReqBody struct {

	// 是否支持容灾回切(仅支持数据库版本大于等于3.100)   - true支持   - false不支持(默认false)
	IsSupportRestore *bool `json:"is_support_restore,omitempty"`

	// 容灾类型。
	DisasterType DisasterFailoverReqBodyDisasterType `json:"disaster_type"`
}

func (o DisasterFailoverReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisasterFailoverReqBody struct{}"
	}

	return strings.Join([]string{"DisasterFailoverReqBody", string(data)}, " ")
}

type DisasterFailoverReqBodyDisasterType struct {
	value string
}

type DisasterFailoverReqBodyDisasterTypeEnum struct {
	STREAM DisasterFailoverReqBodyDisasterType
	DORADO DisasterFailoverReqBodyDisasterType
}

func GetDisasterFailoverReqBodyDisasterTypeEnum() DisasterFailoverReqBodyDisasterTypeEnum {
	return DisasterFailoverReqBodyDisasterTypeEnum{
		STREAM: DisasterFailoverReqBodyDisasterType{
			value: "stream",
		},
		DORADO: DisasterFailoverReqBodyDisasterType{
			value: "dorado",
		},
	}
}

func (c DisasterFailoverReqBodyDisasterType) Value() string {
	return c.value
}

func (c DisasterFailoverReqBodyDisasterType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DisasterFailoverReqBodyDisasterType) UnmarshalJSON(b []byte) error {
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
