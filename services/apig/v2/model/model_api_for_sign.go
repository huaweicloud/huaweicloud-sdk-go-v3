package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ApiForSign struct {

	// API的认证方式
	AuthType *string `json:"auth_type,omitempty"`

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

	// API类型
	Type *int32 `json:"type,omitempty"`

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
