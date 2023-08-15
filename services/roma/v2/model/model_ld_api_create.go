package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type LdApiCreate struct {

	// 后端API名称。  支持汉字、英文、数字、中划线、下划线、点、斜杠、中英文格式下的小括号和冒号、中文格式下的顿号，且只能以英文、汉字和数字开头。
	Name string `json:"name"`

	// 后端API请求路径。  支持英文、数字、中划线、下划线、点等，且以斜杠（/）开头。
	Path string `json:"path"`

	// 后端API请求方法。  支持GET、PUT、POST、DELETE
	Method LdApiCreateMethod `json:"method"`

	// 后端API描述。  不支持<，>字符
	Description *string `json:"description,omitempty"`

	// 后端API版本  支持英文，数字，下划线，中划线，点。
	Version string `json:"version"`

	// 后端API返回类型
	ContentType LdApiCreateContentType `json:"content_type"`

	// 后端API为签名认证时绑定的签名密钥编号
	ApiSignatureId *string `json:"api_signature_id,omitempty"`

	// 后端API归属的集成应用编号
	RomaAppId string `json:"roma_app_id"`

	// API响应信息是否格式化  true： 对响应信息进行格式化  false：对响应信息格式化不进行格式化
	ReturnFormat *bool `json:"return_format,omitempty"`

	// 后端API的请求参数列表
	Parameters *[]LdApiParameter `json:"parameters,omitempty"`
}

func (o LdApiCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LdApiCreate struct{}"
	}

	return strings.Join([]string{"LdApiCreate", string(data)}, " ")
}

type LdApiCreateMethod struct {
	value string
}

type LdApiCreateMethodEnum struct {
	GET    LdApiCreateMethod
	PUT    LdApiCreateMethod
	POST   LdApiCreateMethod
	DELETE LdApiCreateMethod
}

func GetLdApiCreateMethodEnum() LdApiCreateMethodEnum {
	return LdApiCreateMethodEnum{
		GET: LdApiCreateMethod{
			value: "GET",
		},
		PUT: LdApiCreateMethod{
			value: "PUT",
		},
		POST: LdApiCreateMethod{
			value: "POST",
		},
		DELETE: LdApiCreateMethod{
			value: "DELETE",
		},
	}
}

func (c LdApiCreateMethod) Value() string {
	return c.value
}

func (c LdApiCreateMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LdApiCreateMethod) UnmarshalJSON(b []byte) error {
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

type LdApiCreateContentType struct {
	value string
}

type LdApiCreateContentTypeEnum struct {
	JSON   LdApiCreateContentType
	XML    LdApiCreateContentType
	STREAM LdApiCreateContentType
}

func GetLdApiCreateContentTypeEnum() LdApiCreateContentTypeEnum {
	return LdApiCreateContentTypeEnum{
		JSON: LdApiCreateContentType{
			value: "json",
		},
		XML: LdApiCreateContentType{
			value: "xml",
		},
		STREAM: LdApiCreateContentType{
			value: "stream",
		},
	}
}

func (c LdApiCreateContentType) Value() string {
	return c.value
}

func (c LdApiCreateContentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LdApiCreateContentType) UnmarshalJSON(b []byte) error {
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
