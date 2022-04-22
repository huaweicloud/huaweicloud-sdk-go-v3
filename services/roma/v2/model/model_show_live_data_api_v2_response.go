package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// Response Object
type ShowLiveDataApiV2Response struct {

	// 后端API名称。  支持汉字、英文、数字、中划线、下划线、点、斜杠、中英文格式下的小括号和冒号、中文格式下的顿号，且只能以英文、汉字和数字开头。
	Name string `json:"name"`

	// 后端API请求路径。  支持英文、数字、中划线、下划线、点等，且以斜杠（/）开头。
	Path string `json:"path"`

	// 后端API请求方法。  支持GET、PUT、POST、DELETE
	Method ShowLiveDataApiV2ResponseMethod `json:"method"`

	// 后端API描述。  不支持<，>字符
	Description *string `json:"description,omitempty"`

	// 后端API版本  支持英文，数字，下划线，中划线，点。
	Version string `json:"version"`

	// 后端API返回类型
	ContentType ShowLiveDataApiV2ResponseContentType `json:"content_type"`

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
	Type *ShowLiveDataApiV2ResponseType `json:"type,omitempty"`

	// 后端API状态： - 1：待开发 - 3：开发中 - 4：已部署
	Status *ShowLiveDataApiV2ResponseStatus `json:"status,omitempty"`

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

func (o ShowLiveDataApiV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLiveDataApiV2Response struct{}"
	}

	return strings.Join([]string{"ShowLiveDataApiV2Response", string(data)}, " ")
}

type ShowLiveDataApiV2ResponseMethod struct {
	value string
}

type ShowLiveDataApiV2ResponseMethodEnum struct {
	GET    ShowLiveDataApiV2ResponseMethod
	PUT    ShowLiveDataApiV2ResponseMethod
	POST   ShowLiveDataApiV2ResponseMethod
	DELETE ShowLiveDataApiV2ResponseMethod
}

func GetShowLiveDataApiV2ResponseMethodEnum() ShowLiveDataApiV2ResponseMethodEnum {
	return ShowLiveDataApiV2ResponseMethodEnum{
		GET: ShowLiveDataApiV2ResponseMethod{
			value: "GET",
		},
		PUT: ShowLiveDataApiV2ResponseMethod{
			value: "PUT",
		},
		POST: ShowLiveDataApiV2ResponseMethod{
			value: "POST",
		},
		DELETE: ShowLiveDataApiV2ResponseMethod{
			value: "DELETE",
		},
	}
}

func (c ShowLiveDataApiV2ResponseMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowLiveDataApiV2ResponseMethod) UnmarshalJSON(b []byte) error {
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

type ShowLiveDataApiV2ResponseContentType struct {
	value string
}

type ShowLiveDataApiV2ResponseContentTypeEnum struct {
	JSON   ShowLiveDataApiV2ResponseContentType
	XML    ShowLiveDataApiV2ResponseContentType
	STREAM ShowLiveDataApiV2ResponseContentType
}

func GetShowLiveDataApiV2ResponseContentTypeEnum() ShowLiveDataApiV2ResponseContentTypeEnum {
	return ShowLiveDataApiV2ResponseContentTypeEnum{
		JSON: ShowLiveDataApiV2ResponseContentType{
			value: "json",
		},
		XML: ShowLiveDataApiV2ResponseContentType{
			value: "xml",
		},
		STREAM: ShowLiveDataApiV2ResponseContentType{
			value: "stream",
		},
	}
}

func (c ShowLiveDataApiV2ResponseContentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowLiveDataApiV2ResponseContentType) UnmarshalJSON(b []byte) error {
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

type ShowLiveDataApiV2ResponseType struct {
	value string
}

type ShowLiveDataApiV2ResponseTypeEnum struct {
	DATA     ShowLiveDataApiV2ResponseType
	FUNCTION ShowLiveDataApiV2ResponseType
}

func GetShowLiveDataApiV2ResponseTypeEnum() ShowLiveDataApiV2ResponseTypeEnum {
	return ShowLiveDataApiV2ResponseTypeEnum{
		DATA: ShowLiveDataApiV2ResponseType{
			value: "data",
		},
		FUNCTION: ShowLiveDataApiV2ResponseType{
			value: "function",
		},
	}
}

func (c ShowLiveDataApiV2ResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowLiveDataApiV2ResponseType) UnmarshalJSON(b []byte) error {
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

type ShowLiveDataApiV2ResponseStatus struct {
	value int32
}

type ShowLiveDataApiV2ResponseStatusEnum struct {
	E_1 ShowLiveDataApiV2ResponseStatus
	E_3 ShowLiveDataApiV2ResponseStatus
	E_4 ShowLiveDataApiV2ResponseStatus
}

func GetShowLiveDataApiV2ResponseStatusEnum() ShowLiveDataApiV2ResponseStatusEnum {
	return ShowLiveDataApiV2ResponseStatusEnum{
		E_1: ShowLiveDataApiV2ResponseStatus{
			value: 1,
		}, E_3: ShowLiveDataApiV2ResponseStatus{
			value: 3,
		}, E_4: ShowLiveDataApiV2ResponseStatus{
			value: 4,
		},
	}
}

func (c ShowLiveDataApiV2ResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowLiveDataApiV2ResponseStatus) UnmarshalJSON(b []byte) error {
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
