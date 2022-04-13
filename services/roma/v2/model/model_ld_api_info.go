package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"errors"
	"strings"
)

// 后端API详情
type LdApiInfo struct {
	// 后端API名称。  支持汉字、英文、数字、中划线、下划线、点、斜杠、中英文格式下的小括号和冒号、中文格式下的顿号，且只能以英文、汉字和数字开头。

	Name string `json:"name"`
	// 后端API请求路径。  支持英文、数字、中划线、下划线、点等，且以斜杠（/）开头。

	Path string `json:"path"`
	// 后端API请求方法。  支持GET、PUT、POST、DELETE

	Method LdApiInfoMethod `json:"method"`
	// 后端API描述。  不支持<，>字符

	Description *string `json:"description,omitempty"`
	// 后端API版本  支持英文，数字，下划线，中划线，点。

	Version string `json:"version"`
	// 后端API返回类型

	ContentType LdApiInfoContentType `json:"content_type"`
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

	Type *LdApiInfoType `json:"type,omitempty"`
	// 后端API状态： - 1：待开发 - 3：开发中 - 4：已部署

	Status *LdApiInfoStatus `json:"status,omitempty"`
	// 后端API创建时间

	CreatedTime *sdktime.SdkTime `json:"created_time,omitempty"`
	// 后端API修改时间

	ModifiedTime *sdktime.SdkTime `json:"modified_time,omitempty"`
	// 后端API脚本信息

	Scripts *[]LdApiScript `json:"scripts,omitempty"`
	// 后端API归属的集成应用名称

	RomaAppName *string `json:"roma_app_name,omitempty"`
}

func (o LdApiInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LdApiInfo struct{}"
	}

	return strings.Join([]string{"LdApiInfo", string(data)}, " ")
}

type LdApiInfoMethod struct {
	value string
}

type LdApiInfoMethodEnum struct {
	GET    LdApiInfoMethod
	PUT    LdApiInfoMethod
	POST   LdApiInfoMethod
	DELETE LdApiInfoMethod
}

func GetLdApiInfoMethodEnum() LdApiInfoMethodEnum {
	return LdApiInfoMethodEnum{
		GET: LdApiInfoMethod{
			value: "GET",
		},
		PUT: LdApiInfoMethod{
			value: "PUT",
		},
		POST: LdApiInfoMethod{
			value: "POST",
		},
		DELETE: LdApiInfoMethod{
			value: "DELETE",
		},
	}
}

func (c LdApiInfoMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LdApiInfoMethod) UnmarshalJSON(b []byte) error {
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

type LdApiInfoContentType struct {
	value string
}

type LdApiInfoContentTypeEnum struct {
	JSON   LdApiInfoContentType
	XML    LdApiInfoContentType
	STREAM LdApiInfoContentType
}

func GetLdApiInfoContentTypeEnum() LdApiInfoContentTypeEnum {
	return LdApiInfoContentTypeEnum{
		JSON: LdApiInfoContentType{
			value: "json",
		},
		XML: LdApiInfoContentType{
			value: "xml",
		},
		STREAM: LdApiInfoContentType{
			value: "stream",
		},
	}
}

func (c LdApiInfoContentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LdApiInfoContentType) UnmarshalJSON(b []byte) error {
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

type LdApiInfoType struct {
	value string
}

type LdApiInfoTypeEnum struct {
	DATA     LdApiInfoType
	FUNCTION LdApiInfoType
}

func GetLdApiInfoTypeEnum() LdApiInfoTypeEnum {
	return LdApiInfoTypeEnum{
		DATA: LdApiInfoType{
			value: "data",
		},
		FUNCTION: LdApiInfoType{
			value: "function",
		},
	}
}

func (c LdApiInfoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LdApiInfoType) UnmarshalJSON(b []byte) error {
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

type LdApiInfoStatus struct {
	value int32
}

type LdApiInfoStatusEnum struct {
	E_1 LdApiInfoStatus
	E_3 LdApiInfoStatus
	E_4 LdApiInfoStatus
}

func GetLdApiInfoStatusEnum() LdApiInfoStatusEnum {
	return LdApiInfoStatusEnum{
		E_1: LdApiInfoStatus{
			value: 1,
		}, E_3: LdApiInfoStatus{
			value: 3,
		}, E_4: LdApiInfoStatus{
			value: 4,
		},
	}
}

func (c LdApiInfoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LdApiInfoStatus) UnmarshalJSON(b []byte) error {
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
