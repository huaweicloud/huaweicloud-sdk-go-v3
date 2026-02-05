package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ApiForSign struct {

	// API的认证方式。 - NONE：无认证 - APP：APP认证 - IAM：IAM认证 - AUTHORIZER：自定义认证
	AuthType *ApiForSignAuthType `json:"auth_type,omitempty"`

	// 发布的环境名
	RunEnvName *string `json:"run_env_name,omitempty"`

	// API所属分组的名称
	GroupName *string `json:"group_name,omitempty"`

	// 发布记录的编号
	PublishId *string `json:"publish_id,omitempty"`

	// API所属分组的编号
	GroupId *string `json:"group_id,omitempty"`

	// API名称
	Name *string `json:"name,omitempty"`

	// API描述
	Remark *string `json:"remark,omitempty"`

	// 发布的环境id
	RunEnvId *string `json:"run_env_id,omitempty"`

	// API编号
	Id *string `json:"id,omitempty"`

	// API的请求地址
	ReqUri *string `json:"req_uri,omitempty"`

	// API绑定的标签，标签配额默认10条，可以联系技术调整。
	Tags *[]string `json:"tags,omitempty"`

	// API类型。 - 1：公有API - 2：私有API
	Type *ApiForSignType `json:"type,omitempty"`

	// 已绑定的签名密钥名称
	SignatureName *string `json:"signature_name,omitempty"`

	// API请求方法
	ReqMethod *ApiForSignReqMethod `json:"req_method,omitempty"`
}

func (o ApiForSign) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiForSign struct{}"
	}

	return strings.Join([]string{"ApiForSign", string(data)}, " ")
}

type ApiForSignAuthType struct {
	value string
}

type ApiForSignAuthTypeEnum struct {
	NONE       ApiForSignAuthType
	APP        ApiForSignAuthType
	IAM        ApiForSignAuthType
	AUTHORIZER ApiForSignAuthType
}

func GetApiForSignAuthTypeEnum() ApiForSignAuthTypeEnum {
	return ApiForSignAuthTypeEnum{
		NONE: ApiForSignAuthType{
			value: "NONE",
		},
		APP: ApiForSignAuthType{
			value: "APP",
		},
		IAM: ApiForSignAuthType{
			value: "IAM",
		},
		AUTHORIZER: ApiForSignAuthType{
			value: "AUTHORIZER",
		},
	}
}

func (c ApiForSignAuthType) Value() string {
	return c.value
}

func (c ApiForSignAuthType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiForSignAuthType) UnmarshalJSON(b []byte) error {
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

type ApiForSignType struct {
	value int32
}

type ApiForSignTypeEnum struct {
	E_1 ApiForSignType
	E_2 ApiForSignType
}

func GetApiForSignTypeEnum() ApiForSignTypeEnum {
	return ApiForSignTypeEnum{
		E_1: ApiForSignType{
			value: 1,
		}, E_2: ApiForSignType{
			value: 2,
		},
	}
}

func (c ApiForSignType) Value() int32 {
	return c.value
}

func (c ApiForSignType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiForSignType) UnmarshalJSON(b []byte) error {
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

type ApiForSignReqMethod struct {
	value string
}

type ApiForSignReqMethodEnum struct {
	GET     ApiForSignReqMethod
	POST    ApiForSignReqMethod
	DELETE  ApiForSignReqMethod
	PUT     ApiForSignReqMethod
	PATCH   ApiForSignReqMethod
	HEAD    ApiForSignReqMethod
	OPTIONS ApiForSignReqMethod
	ANY     ApiForSignReqMethod
}

func GetApiForSignReqMethodEnum() ApiForSignReqMethodEnum {
	return ApiForSignReqMethodEnum{
		GET: ApiForSignReqMethod{
			value: "GET",
		},
		POST: ApiForSignReqMethod{
			value: "POST",
		},
		DELETE: ApiForSignReqMethod{
			value: "DELETE",
		},
		PUT: ApiForSignReqMethod{
			value: "PUT",
		},
		PATCH: ApiForSignReqMethod{
			value: "PATCH",
		},
		HEAD: ApiForSignReqMethod{
			value: "HEAD",
		},
		OPTIONS: ApiForSignReqMethod{
			value: "OPTIONS",
		},
		ANY: ApiForSignReqMethod{
			value: "ANY",
		},
	}
}

func (c ApiForSignReqMethod) Value() string {
	return c.value
}

func (c ApiForSignReqMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiForSignReqMethod) UnmarshalJSON(b []byte) error {
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
