package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// Response Object
type CreateLiveDataApiV2Response struct {
	// 后端API名称。  支持汉字、英文、数字、中划线、下划线、点、斜杠、中英文格式下的小括号和冒号、中文格式下的顿号，且只能以英文、汉字和数字开头。

	Name string `json:"name"`
	// 后端API请求路径。  支持英文、数字、中划线、下划线、点等，且以斜杠（/）开头。

	Path string `json:"path"`
	// 后端API请求方法。  支持GET、PUT、POST、DELETE

	Method CreateLiveDataApiV2ResponseMethod `json:"method"`
	// 后端API描述。  不支持<，>字符

	Description *string `json:"description,omitempty"`
	// 后端API版本  支持英文，数字，下划线，中划线，点。

	Version string `json:"version"`
	// 后端API返回类型

	ContentType CreateLiveDataApiV2ResponseContentType `json:"content_type"`
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

	Type *CreateLiveDataApiV2ResponseType `json:"type,omitempty"`
	// 后端API状态： - 1：待开发 - 3：开发中 - 4：已部署

	Status *CreateLiveDataApiV2ResponseStatus `json:"status,omitempty"`
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

func (o CreateLiveDataApiV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLiveDataApiV2Response struct{}"
	}

	return strings.Join([]string{"CreateLiveDataApiV2Response", string(data)}, " ")
}

type CreateLiveDataApiV2ResponseMethod struct {
	value string
}

type CreateLiveDataApiV2ResponseMethodEnum struct {
	GET    CreateLiveDataApiV2ResponseMethod
	PUT    CreateLiveDataApiV2ResponseMethod
	POST   CreateLiveDataApiV2ResponseMethod
	DELETE CreateLiveDataApiV2ResponseMethod
}

func GetCreateLiveDataApiV2ResponseMethodEnum() CreateLiveDataApiV2ResponseMethodEnum {
	return CreateLiveDataApiV2ResponseMethodEnum{
		GET: CreateLiveDataApiV2ResponseMethod{
			value: "GET",
		},
		PUT: CreateLiveDataApiV2ResponseMethod{
			value: "PUT",
		},
		POST: CreateLiveDataApiV2ResponseMethod{
			value: "POST",
		},
		DELETE: CreateLiveDataApiV2ResponseMethod{
			value: "DELETE",
		},
	}
}

func (c CreateLiveDataApiV2ResponseMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateLiveDataApiV2ResponseMethod) UnmarshalJSON(b []byte) error {
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

type CreateLiveDataApiV2ResponseContentType struct {
	value string
}

type CreateLiveDataApiV2ResponseContentTypeEnum struct {
	JSON   CreateLiveDataApiV2ResponseContentType
	XML    CreateLiveDataApiV2ResponseContentType
	STREAM CreateLiveDataApiV2ResponseContentType
}

func GetCreateLiveDataApiV2ResponseContentTypeEnum() CreateLiveDataApiV2ResponseContentTypeEnum {
	return CreateLiveDataApiV2ResponseContentTypeEnum{
		JSON: CreateLiveDataApiV2ResponseContentType{
			value: "json",
		},
		XML: CreateLiveDataApiV2ResponseContentType{
			value: "xml",
		},
		STREAM: CreateLiveDataApiV2ResponseContentType{
			value: "stream",
		},
	}
}

func (c CreateLiveDataApiV2ResponseContentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateLiveDataApiV2ResponseContentType) UnmarshalJSON(b []byte) error {
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

type CreateLiveDataApiV2ResponseType struct {
	value string
}

type CreateLiveDataApiV2ResponseTypeEnum struct {
	DATA     CreateLiveDataApiV2ResponseType
	FUNCTION CreateLiveDataApiV2ResponseType
}

func GetCreateLiveDataApiV2ResponseTypeEnum() CreateLiveDataApiV2ResponseTypeEnum {
	return CreateLiveDataApiV2ResponseTypeEnum{
		DATA: CreateLiveDataApiV2ResponseType{
			value: "data",
		},
		FUNCTION: CreateLiveDataApiV2ResponseType{
			value: "function",
		},
	}
}

func (c CreateLiveDataApiV2ResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateLiveDataApiV2ResponseType) UnmarshalJSON(b []byte) error {
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

type CreateLiveDataApiV2ResponseStatus struct {
	value int32
}

type CreateLiveDataApiV2ResponseStatusEnum struct {
	E_1 CreateLiveDataApiV2ResponseStatus
	E_3 CreateLiveDataApiV2ResponseStatus
	E_4 CreateLiveDataApiV2ResponseStatus
}

func GetCreateLiveDataApiV2ResponseStatusEnum() CreateLiveDataApiV2ResponseStatusEnum {
	return CreateLiveDataApiV2ResponseStatusEnum{
		E_1: CreateLiveDataApiV2ResponseStatus{
			value: 1,
		}, E_3: CreateLiveDataApiV2ResponseStatus{
			value: 3,
		}, E_4: CreateLiveDataApiV2ResponseStatus{
			value: 4,
		},
	}
}

func (c CreateLiveDataApiV2ResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateLiveDataApiV2ResponseStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
