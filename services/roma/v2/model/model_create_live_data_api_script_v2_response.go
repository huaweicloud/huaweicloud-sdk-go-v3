package model

import (
	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// CreateLiveDataApiScriptV2Response Response Object
type CreateLiveDataApiScriptV2Response struct {

	// 后端API名称。  支持汉字、英文、数字、中划线、下划线、点、斜杠、中英文格式下的小括号和冒号、中文格式下的顿号，且只能以英文、汉字和数字开头。
	Name string `json:"name"`

	// 后端API请求路径。  支持英文、数字、中划线、下划线、点等，且以斜杠（/）开头。
	Path string `json:"path"`

	// 后端API请求方法。  支持GET、PUT、POST、DELETE
	Method CreateLiveDataApiScriptV2ResponseMethod `json:"method"`

	// 后端API描述。  不支持<，>字符
	Description *string `json:"description,omitempty"`

	// 后端API版本  支持英文，数字，下划线，中划线，点。
	Version string `json:"version"`

	// 后端API返回类型
	ContentType CreateLiveDataApiScriptV2ResponseContentType `json:"content_type"`

	// 后端API为签名认证时绑定的签名密钥编号
	ApiSignatureId *string `json:"api_signature_id,omitempty"`

	// 后端API归属的集成应用编号
	RomaAppId string `json:"roma_app_id"`

	// API响应信息是否格式化  true： 对响应信息进行格式化  false：对响应信息格式化不进行格式化
	ReturnFormat *bool `json:"return_format,omitempty"`

	// 后端API的请求参数列表
	Parameters *[]LdApiParameter `json:"parameters,omitempty"`

	// 后端API编号
	Id *string `json:"id,omitempty"`

	// 后端API所属实例编号
	Instance *string `json:"instance,omitempty"`

	// 后端API类型： - data：数据后端 - function： 函数后端
	Type *CreateLiveDataApiScriptV2ResponseType `json:"type,omitempty"`

	// 后端API状态： - 1：待开发 - 3：开发中 - 4：已部署
	Status *CreateLiveDataApiScriptV2ResponseStatus `json:"status,omitempty"`

	// 后端API创建时间
	CreatedTime *sdktime.SdkTime `json:"created_time,omitempty"`

	// 后端API修改时间
	ModifiedTime *sdktime.SdkTime `json:"modified_time,omitempty"`

	// 后端API脚本信息
	Scripts *[]LdApiScript `json:"scripts,omitempty"`

	// 后端API归属的集成应用名称
	RomaAppName    *string `json:"roma_app_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateLiveDataApiScriptV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLiveDataApiScriptV2Response struct{}"
	}

	return strings.Join([]string{"CreateLiveDataApiScriptV2Response", string(data)}, " ")
}

type CreateLiveDataApiScriptV2ResponseMethod struct {
	value string
}

type CreateLiveDataApiScriptV2ResponseMethodEnum struct {
	GET    CreateLiveDataApiScriptV2ResponseMethod
	PUT    CreateLiveDataApiScriptV2ResponseMethod
	POST   CreateLiveDataApiScriptV2ResponseMethod
	DELETE CreateLiveDataApiScriptV2ResponseMethod
}

func GetCreateLiveDataApiScriptV2ResponseMethodEnum() CreateLiveDataApiScriptV2ResponseMethodEnum {
	return CreateLiveDataApiScriptV2ResponseMethodEnum{
		GET: CreateLiveDataApiScriptV2ResponseMethod{
			value: "GET",
		},
		PUT: CreateLiveDataApiScriptV2ResponseMethod{
			value: "PUT",
		},
		POST: CreateLiveDataApiScriptV2ResponseMethod{
			value: "POST",
		},
		DELETE: CreateLiveDataApiScriptV2ResponseMethod{
			value: "DELETE",
		},
	}
}

func (c CreateLiveDataApiScriptV2ResponseMethod) Value() string {
	return c.value
}

func (c CreateLiveDataApiScriptV2ResponseMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateLiveDataApiScriptV2ResponseMethod) UnmarshalJSON(b []byte) error {
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

type CreateLiveDataApiScriptV2ResponseContentType struct {
	value string
}

type CreateLiveDataApiScriptV2ResponseContentTypeEnum struct {
	JSON   CreateLiveDataApiScriptV2ResponseContentType
	XML    CreateLiveDataApiScriptV2ResponseContentType
	STREAM CreateLiveDataApiScriptV2ResponseContentType
}

func GetCreateLiveDataApiScriptV2ResponseContentTypeEnum() CreateLiveDataApiScriptV2ResponseContentTypeEnum {
	return CreateLiveDataApiScriptV2ResponseContentTypeEnum{
		JSON: CreateLiveDataApiScriptV2ResponseContentType{
			value: "json",
		},
		XML: CreateLiveDataApiScriptV2ResponseContentType{
			value: "xml",
		},
		STREAM: CreateLiveDataApiScriptV2ResponseContentType{
			value: "stream",
		},
	}
}

func (c CreateLiveDataApiScriptV2ResponseContentType) Value() string {
	return c.value
}

func (c CreateLiveDataApiScriptV2ResponseContentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateLiveDataApiScriptV2ResponseContentType) UnmarshalJSON(b []byte) error {
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

type CreateLiveDataApiScriptV2ResponseType struct {
	value string
}

type CreateLiveDataApiScriptV2ResponseTypeEnum struct {
	DATA     CreateLiveDataApiScriptV2ResponseType
	FUNCTION CreateLiveDataApiScriptV2ResponseType
}

func GetCreateLiveDataApiScriptV2ResponseTypeEnum() CreateLiveDataApiScriptV2ResponseTypeEnum {
	return CreateLiveDataApiScriptV2ResponseTypeEnum{
		DATA: CreateLiveDataApiScriptV2ResponseType{
			value: "data",
		},
		FUNCTION: CreateLiveDataApiScriptV2ResponseType{
			value: "function",
		},
	}
}

func (c CreateLiveDataApiScriptV2ResponseType) Value() string {
	return c.value
}

func (c CreateLiveDataApiScriptV2ResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateLiveDataApiScriptV2ResponseType) UnmarshalJSON(b []byte) error {
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

type CreateLiveDataApiScriptV2ResponseStatus struct {
	value int32
}

type CreateLiveDataApiScriptV2ResponseStatusEnum struct {
	E_1 CreateLiveDataApiScriptV2ResponseStatus
	E_3 CreateLiveDataApiScriptV2ResponseStatus
	E_4 CreateLiveDataApiScriptV2ResponseStatus
}

func GetCreateLiveDataApiScriptV2ResponseStatusEnum() CreateLiveDataApiScriptV2ResponseStatusEnum {
	return CreateLiveDataApiScriptV2ResponseStatusEnum{
		E_1: CreateLiveDataApiScriptV2ResponseStatus{
			value: 1,
		}, E_3: CreateLiveDataApiScriptV2ResponseStatus{
			value: 3,
		}, E_4: CreateLiveDataApiScriptV2ResponseStatus{
			value: 4,
		},
	}
}

func (c CreateLiveDataApiScriptV2ResponseStatus) Value() int32 {
	return c.value
}

func (c CreateLiveDataApiScriptV2ResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateLiveDataApiScriptV2ResponseStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
