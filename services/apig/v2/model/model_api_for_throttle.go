package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type ApiForThrottle struct {

	// API的认证方式。 - NONE：无认证 - APP：APP认证 - IAM：IAM认证 - AUTHORIZER：自定义认证
	AuthType *ApiForThrottleAuthType `json:"auth_type,omitempty"`

	// API所属分组的名称
	GroupName *string `json:"group_name,omitempty"`

	// API的发布记录编号
	PublishId *string `json:"publish_id,omitempty"`

	// 与流控策略的绑定关系编号
	ThrottleApplyId *string `json:"throttle_apply_id,omitempty"`

	// 已绑定的流控策略的绑定时间
	ApplyTime *sdktime.SdkTime `json:"apply_time,omitempty"`

	// API描述
	Remark *string `json:"remark,omitempty"`

	// 发布的环境id
	RunEnvId *string `json:"run_env_id,omitempty"`

	// API类型。 - 1：公有API - 2：私有API
	Type *ApiForThrottleType `json:"type,omitempty"`

	// 绑定的流控策略名称
	ThrottleName *string `json:"throttle_name,omitempty"`

	// API的访问地址
	ReqUri *string `json:"req_uri,omitempty"`

	// 发布的环境名
	RunEnvName *string `json:"run_env_name,omitempty"`

	// API所属分组的编号
	GroupId *string `json:"group_id,omitempty"`

	// API名称
	Name *string `json:"name,omitempty"`

	// API编号
	Id *string `json:"id,omitempty"`

	// API请求方法
	ReqMethod *ApiForThrottleReqMethod `json:"req_method,omitempty"`

	// API绑定的标签，标签配额默认10条，可以联系技术调整。
	Tags *[]string `json:"tags,omitempty"`
}

func (o ApiForThrottle) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiForThrottle struct{}"
	}

	return strings.Join([]string{"ApiForThrottle", string(data)}, " ")
}

type ApiForThrottleAuthType struct {
	value string
}

type ApiForThrottleAuthTypeEnum struct {
	NONE       ApiForThrottleAuthType
	APP        ApiForThrottleAuthType
	IAM        ApiForThrottleAuthType
	AUTHORIZER ApiForThrottleAuthType
}

func GetApiForThrottleAuthTypeEnum() ApiForThrottleAuthTypeEnum {
	return ApiForThrottleAuthTypeEnum{
		NONE: ApiForThrottleAuthType{
			value: "NONE",
		},
		APP: ApiForThrottleAuthType{
			value: "APP",
		},
		IAM: ApiForThrottleAuthType{
			value: "IAM",
		},
		AUTHORIZER: ApiForThrottleAuthType{
			value: "AUTHORIZER",
		},
	}
}

func (c ApiForThrottleAuthType) Value() string {
	return c.value
}

func (c ApiForThrottleAuthType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiForThrottleAuthType) UnmarshalJSON(b []byte) error {
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

type ApiForThrottleType struct {
	value int32
}

type ApiForThrottleTypeEnum struct {
	E_1 ApiForThrottleType
	E_2 ApiForThrottleType
}

func GetApiForThrottleTypeEnum() ApiForThrottleTypeEnum {
	return ApiForThrottleTypeEnum{
		E_1: ApiForThrottleType{
			value: 1,
		}, E_2: ApiForThrottleType{
			value: 2,
		},
	}
}

func (c ApiForThrottleType) Value() int32 {
	return c.value
}

func (c ApiForThrottleType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiForThrottleType) UnmarshalJSON(b []byte) error {
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

type ApiForThrottleReqMethod struct {
	value string
}

type ApiForThrottleReqMethodEnum struct {
	GET     ApiForThrottleReqMethod
	POST    ApiForThrottleReqMethod
	DELETE  ApiForThrottleReqMethod
	PUT     ApiForThrottleReqMethod
	PATCH   ApiForThrottleReqMethod
	HEAD    ApiForThrottleReqMethod
	OPTIONS ApiForThrottleReqMethod
	ANY     ApiForThrottleReqMethod
}

func GetApiForThrottleReqMethodEnum() ApiForThrottleReqMethodEnum {
	return ApiForThrottleReqMethodEnum{
		GET: ApiForThrottleReqMethod{
			value: "GET",
		},
		POST: ApiForThrottleReqMethod{
			value: "POST",
		},
		DELETE: ApiForThrottleReqMethod{
			value: "DELETE",
		},
		PUT: ApiForThrottleReqMethod{
			value: "PUT",
		},
		PATCH: ApiForThrottleReqMethod{
			value: "PATCH",
		},
		HEAD: ApiForThrottleReqMethod{
			value: "HEAD",
		},
		OPTIONS: ApiForThrottleReqMethod{
			value: "OPTIONS",
		},
		ANY: ApiForThrottleReqMethod{
			value: "ANY",
		},
	}
}

func (c ApiForThrottleReqMethod) Value() string {
	return c.value
}

func (c ApiForThrottleReqMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiForThrottleReqMethod) UnmarshalJSON(b []byte) error {
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
