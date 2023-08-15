package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// MicroserviceApiCreate 导入微服务创建单个API的对象
type MicroserviceApiCreate struct {

	// API名称。  支持汉字、英文、数字、中划线、下划线、点、斜杠、中英文格式下的小括号和冒号、中文格式下的顿号，且只能以英文、汉字和数字开头。 > 中文字符必须为UTF-8或者unicode编码。
	Name *string `json:"name,omitempty"`

	// API的请求方式
	ReqMethod *MicroserviceApiCreateReqMethod `json:"req_method,omitempty"`

	// 请求地址。可以包含请求参数，用{}标识，比如/getUserInfo/{userId}，支持 * % - _ . 等特殊字符，总长度不超过512，且满足URI规范。  /apic/health_check为APIG预置的健康检查路径，当req_method=GET时不支持req_uri=/apic/health_check。  > 需要服从URI规范。
	ReqUri string `json:"req_uri"`

	// API的匹配方式 - SWA：前缀匹配 - NORMAL：正常匹配（绝对匹配） 默认：NORMAL
	MatchMode *MicroserviceApiCreateMatchMode `json:"match_mode,omitempty"`
}

func (o MicroserviceApiCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MicroserviceApiCreate struct{}"
	}

	return strings.Join([]string{"MicroserviceApiCreate", string(data)}, " ")
}

type MicroserviceApiCreateReqMethod struct {
	value string
}

type MicroserviceApiCreateReqMethodEnum struct {
	GET     MicroserviceApiCreateReqMethod
	POST    MicroserviceApiCreateReqMethod
	PUT     MicroserviceApiCreateReqMethod
	DELETE  MicroserviceApiCreateReqMethod
	HEAD    MicroserviceApiCreateReqMethod
	PATCH   MicroserviceApiCreateReqMethod
	OPTIONS MicroserviceApiCreateReqMethod
	ANY     MicroserviceApiCreateReqMethod
}

func GetMicroserviceApiCreateReqMethodEnum() MicroserviceApiCreateReqMethodEnum {
	return MicroserviceApiCreateReqMethodEnum{
		GET: MicroserviceApiCreateReqMethod{
			value: "GET",
		},
		POST: MicroserviceApiCreateReqMethod{
			value: "POST",
		},
		PUT: MicroserviceApiCreateReqMethod{
			value: "PUT",
		},
		DELETE: MicroserviceApiCreateReqMethod{
			value: "DELETE",
		},
		HEAD: MicroserviceApiCreateReqMethod{
			value: "HEAD",
		},
		PATCH: MicroserviceApiCreateReqMethod{
			value: "PATCH",
		},
		OPTIONS: MicroserviceApiCreateReqMethod{
			value: "OPTIONS",
		},
		ANY: MicroserviceApiCreateReqMethod{
			value: "ANY",
		},
	}
}

func (c MicroserviceApiCreateReqMethod) Value() string {
	return c.value
}

func (c MicroserviceApiCreateReqMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MicroserviceApiCreateReqMethod) UnmarshalJSON(b []byte) error {
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

type MicroserviceApiCreateMatchMode struct {
	value string
}

type MicroserviceApiCreateMatchModeEnum struct {
	SWA    MicroserviceApiCreateMatchMode
	NORMAL MicroserviceApiCreateMatchMode
}

func GetMicroserviceApiCreateMatchModeEnum() MicroserviceApiCreateMatchModeEnum {
	return MicroserviceApiCreateMatchModeEnum{
		SWA: MicroserviceApiCreateMatchMode{
			value: "SWA",
		},
		NORMAL: MicroserviceApiCreateMatchMode{
			value: "NORMAL",
		},
	}
}

func (c MicroserviceApiCreateMatchMode) Value() string {
	return c.value
}

func (c MicroserviceApiCreateMatchMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MicroserviceApiCreateMatchMode) UnmarshalJSON(b []byte) error {
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
