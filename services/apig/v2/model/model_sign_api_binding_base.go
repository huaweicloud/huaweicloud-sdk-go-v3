package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type SignApiBindingBase struct {

	// API的发布编号
	PublishId *string `json:"publish_id,omitempty"`

	// API编号
	ApiId *string `json:"api_id,omitempty"`

	// API所属分组的名称
	GroupName *string `json:"group_name,omitempty"`

	// 绑定时间
	BindingTime *sdktime.SdkTime `json:"binding_time,omitempty"`

	// API所属环境的编号
	EnvId *string `json:"env_id,omitempty"`

	// API所属环境的名称
	EnvName *string `json:"env_name,omitempty"`

	// API类型
	ApiType *int32 `json:"api_type,omitempty"`

	// API名称
	ApiName *string `json:"api_name,omitempty"`

	// 绑定关系的ID
	Id *string `json:"id,omitempty"`

	// API描述
	ApiRemark *string `json:"api_remark,omitempty"`

	// 签名密钥的编号
	SignId *string `json:"sign_id,omitempty"`

	// 签名密钥的名称。支持汉字，英文，数字，下划线，且只能以英文和汉字开头，3 ~ 64字符。 > 中文字符必须为UTF-8或者unicode编码。
	SignName *string `json:"sign_name,omitempty"`

	// API请求方法
	ReqMethod *SignApiBindingBaseReqMethod `json:"req_method,omitempty"`
}

func (o SignApiBindingBase) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SignApiBindingBase struct{}"
	}

	return strings.Join([]string{"SignApiBindingBase", string(data)}, " ")
}

type SignApiBindingBaseReqMethod struct {
	value string
}

type SignApiBindingBaseReqMethodEnum struct {
	GET     SignApiBindingBaseReqMethod
	POST    SignApiBindingBaseReqMethod
	DELETE  SignApiBindingBaseReqMethod
	PUT     SignApiBindingBaseReqMethod
	PATCH   SignApiBindingBaseReqMethod
	HEAD    SignApiBindingBaseReqMethod
	OPTIONS SignApiBindingBaseReqMethod
	ANY     SignApiBindingBaseReqMethod
}

func GetSignApiBindingBaseReqMethodEnum() SignApiBindingBaseReqMethodEnum {
	return SignApiBindingBaseReqMethodEnum{
		GET: SignApiBindingBaseReqMethod{
			value: "GET",
		},
		POST: SignApiBindingBaseReqMethod{
			value: "POST",
		},
		DELETE: SignApiBindingBaseReqMethod{
			value: "DELETE",
		},
		PUT: SignApiBindingBaseReqMethod{
			value: "PUT",
		},
		PATCH: SignApiBindingBaseReqMethod{
			value: "PATCH",
		},
		HEAD: SignApiBindingBaseReqMethod{
			value: "HEAD",
		},
		OPTIONS: SignApiBindingBaseReqMethod{
			value: "OPTIONS",
		},
		ANY: SignApiBindingBaseReqMethod{
			value: "ANY",
		},
	}
}

func (c SignApiBindingBaseReqMethod) Value() string {
	return c.value
}

func (c SignApiBindingBaseReqMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SignApiBindingBaseReqMethod) UnmarshalJSON(b []byte) error {
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
