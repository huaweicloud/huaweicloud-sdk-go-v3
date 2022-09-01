package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// Response Object
type UpdateLiveDataApiV2Response struct {

	// 后端API名称。  支持汉字、英文、数字、中划线、下划线、点、斜杠、中英文格式下的小括号和冒号、中文格式下的顿号，且只能以英文、汉字和数字开头。
	Name string `json:"name" xml:"name"`

	// 后端API请求路径。  支持英文、数字、中划线、下划线、点等，且以斜杠（/）开头。
	Path string `json:"path" xml:"path"`

	// 后端API请求方法。  支持GET、PUT、POST、DELETE
	Method UpdateLiveDataApiV2ResponseMethod `json:"method" xml:"method"`

	// 后端API描述。  不支持<，>字符
	Description *string `json:"description,omitempty" xml:"description"`

	// 后端API版本  支持英文，数字，下划线，中划线，点。
	Version string `json:"version" xml:"version"`

	// 后端API返回类型
	ContentType UpdateLiveDataApiV2ResponseContentType `json:"content_type" xml:"content_type"`

	// 后端API为签名认证时绑定的签名密钥编号
	ApiSignatureId *string `json:"api_signature_id,omitempty" xml:"api_signature_id"`

	// 后端API归属的集成应用编号
	RomaAppId string `json:"roma_app_id" xml:"roma_app_id"`

	// API响应信息是否格式化  true： 对响应信息进行格式化  false：对响应信息格式化不进行格式化
	ReturnFormat *bool `json:"return_format,omitempty" xml:"return_format"`

	// 后端API的请求参数列表
	Parameters *[]LdApiParameter `json:"parameters,omitempty" xml:"parameters"`

	// 后端API编号
	Id *string `json:"id,omitempty" xml:"id"`

	// 后端API所属实例编号
	Instance *string `json:"instance,omitempty" xml:"instance"`

	// 后端API类型： - data：数据后端 - function： 函数后端
	Type *UpdateLiveDataApiV2ResponseType `json:"type,omitempty" xml:"type"`

	// 后端API状态： - 1：待开发 - 3：开发中 - 4：已部署
	Status *UpdateLiveDataApiV2ResponseStatus `json:"status,omitempty" xml:"status"`

	// 后端API创建时间
	CreatedTime *sdktime.SdkTime `json:"created_time,omitempty" xml:"created_time"`

	// 后端API修改时间
	ModifiedTime *sdktime.SdkTime `json:"modified_time,omitempty" xml:"modified_time"`

	// 后端API脚本信息
	Scripts *[]LdApiScript `json:"scripts,omitempty" xml:"scripts"`

	// 后端API归属的集成应用名称
	RomaAppName    *string `json:"roma_app_name,omitempty" xml:"roma_app_name"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateLiveDataApiV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLiveDataApiV2Response struct{}"
	}

	return strings.Join([]string{"UpdateLiveDataApiV2Response", string(data)}, " ")
}

type UpdateLiveDataApiV2ResponseMethod struct {
	value string
}

type UpdateLiveDataApiV2ResponseMethodEnum struct {
	GET    UpdateLiveDataApiV2ResponseMethod
	PUT    UpdateLiveDataApiV2ResponseMethod
	POST   UpdateLiveDataApiV2ResponseMethod
	DELETE UpdateLiveDataApiV2ResponseMethod
}

func GetUpdateLiveDataApiV2ResponseMethodEnum() UpdateLiveDataApiV2ResponseMethodEnum {
	return UpdateLiveDataApiV2ResponseMethodEnum{
		GET: UpdateLiveDataApiV2ResponseMethod{
			value: "GET",
		},
		PUT: UpdateLiveDataApiV2ResponseMethod{
			value: "PUT",
		},
		POST: UpdateLiveDataApiV2ResponseMethod{
			value: "POST",
		},
		DELETE: UpdateLiveDataApiV2ResponseMethod{
			value: "DELETE",
		},
	}
}

func (c UpdateLiveDataApiV2ResponseMethod) Value() string {
	return c.value
}

func (c UpdateLiveDataApiV2ResponseMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateLiveDataApiV2ResponseMethod) UnmarshalJSON(b []byte) error {
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

type UpdateLiveDataApiV2ResponseContentType struct {
	value string
}

type UpdateLiveDataApiV2ResponseContentTypeEnum struct {
	JSON   UpdateLiveDataApiV2ResponseContentType
	XML    UpdateLiveDataApiV2ResponseContentType
	STREAM UpdateLiveDataApiV2ResponseContentType
}

func GetUpdateLiveDataApiV2ResponseContentTypeEnum() UpdateLiveDataApiV2ResponseContentTypeEnum {
	return UpdateLiveDataApiV2ResponseContentTypeEnum{
		JSON: UpdateLiveDataApiV2ResponseContentType{
			value: "json",
		},
		XML: UpdateLiveDataApiV2ResponseContentType{
			value: "xml",
		},
		STREAM: UpdateLiveDataApiV2ResponseContentType{
			value: "stream",
		},
	}
}

func (c UpdateLiveDataApiV2ResponseContentType) Value() string {
	return c.value
}

func (c UpdateLiveDataApiV2ResponseContentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateLiveDataApiV2ResponseContentType) UnmarshalJSON(b []byte) error {
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

type UpdateLiveDataApiV2ResponseType struct {
	value string
}

type UpdateLiveDataApiV2ResponseTypeEnum struct {
	DATA     UpdateLiveDataApiV2ResponseType
	FUNCTION UpdateLiveDataApiV2ResponseType
}

func GetUpdateLiveDataApiV2ResponseTypeEnum() UpdateLiveDataApiV2ResponseTypeEnum {
	return UpdateLiveDataApiV2ResponseTypeEnum{
		DATA: UpdateLiveDataApiV2ResponseType{
			value: "data",
		},
		FUNCTION: UpdateLiveDataApiV2ResponseType{
			value: "function",
		},
	}
}

func (c UpdateLiveDataApiV2ResponseType) Value() string {
	return c.value
}

func (c UpdateLiveDataApiV2ResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateLiveDataApiV2ResponseType) UnmarshalJSON(b []byte) error {
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

type UpdateLiveDataApiV2ResponseStatus struct {
	value int32
}

type UpdateLiveDataApiV2ResponseStatusEnum struct {
	E_1 UpdateLiveDataApiV2ResponseStatus
	E_3 UpdateLiveDataApiV2ResponseStatus
	E_4 UpdateLiveDataApiV2ResponseStatus
}

func GetUpdateLiveDataApiV2ResponseStatusEnum() UpdateLiveDataApiV2ResponseStatusEnum {
	return UpdateLiveDataApiV2ResponseStatusEnum{
		E_1: UpdateLiveDataApiV2ResponseStatus{
			value: 1,
		}, E_3: UpdateLiveDataApiV2ResponseStatus{
			value: 3,
		}, E_4: UpdateLiveDataApiV2ResponseStatus{
			value: 4,
		},
	}
}

func (c UpdateLiveDataApiV2ResponseStatus) Value() int32 {
	return c.value
}

func (c UpdateLiveDataApiV2ResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateLiveDataApiV2ResponseStatus) UnmarshalJSON(b []byte) error {
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
