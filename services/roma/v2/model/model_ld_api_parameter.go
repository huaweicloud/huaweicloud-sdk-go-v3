package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// LdApiParameter 后端API请求参数
type LdApiParameter struct {

	// 参数名称： - 参数位于Headers、 Parameters时，用户自行定义，支持英文、数字、点、中划线、下划线，且需要英文开头，不区分大小写。 - 参数位于Body时候，参数以application/json、application/xml、application/text为名，但实际是以请求body里的键值对作为参数名和参数值，比如请求消息样例，参数名为application/json，参数值为{\\\"table\\\":\\\"apic01\\\",\\\"id\\\":\\\"1\\\"}，后端取table：apic01，id：1这两个键值对作为入参。 - 注意：定义参数不要重名，否则会覆盖掉，当Headers、Parameters重复时候，Parameters会被覆盖，当Parameters和Body里的键值对重复时候，Parameters会被覆盖。
	Name string `json:"name"`

	// 该参数在调用API时候所放的位置： - Headers ：放于请求头 - Parameters ：放于请求参数 - Body：放于请求体
	In LdApiParameterIn `json:"in"`

	// 参数默认值
	Default *string `json:"default,omitempty"`

	// 参数描述  不支持<，>字符
	Description *string `json:"description,omitempty"`

	// 参数是否必须。true：必须，false：不必须
	Required *bool `json:"required,omitempty"`
}

func (o LdApiParameter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LdApiParameter struct{}"
	}

	return strings.Join([]string{"LdApiParameter", string(data)}, " ")
}

type LdApiParameterIn struct {
	value string
}

type LdApiParameterInEnum struct {
	HEADERS    LdApiParameterIn
	PARAMETERS LdApiParameterIn
	BODY       LdApiParameterIn
}

func GetLdApiParameterInEnum() LdApiParameterInEnum {
	return LdApiParameterInEnum{
		HEADERS: LdApiParameterIn{
			value: "Headers",
		},
		PARAMETERS: LdApiParameterIn{
			value: "Parameters",
		},
		BODY: LdApiParameterIn{
			value: "Body",
		},
	}
}

func (c LdApiParameterIn) Value() string {
	return c.value
}

func (c LdApiParameterIn) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LdApiParameterIn) UnmarshalJSON(b []byte) error {
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
