package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// ApiFunc [函数工作流后端详情](tag:hws,hws_hk,hcs,hcs_sm,fcs,g42)[暂不支持](tag:Site)
type ApiFunc struct {

	// 函数URN
	FunctionUrn string `json:"function_urn"`

	// 对接函数的网络架构类型 - V1：非VPC网络架构 - V2：VPC网络架构
	NetworkType *ApiFuncNetworkType `json:"network_type,omitempty"`

	// 描述信息。 > 中文字符必须为UTF-8或者unicode编码。
	Remark *string `json:"remark,omitempty"`

	// 调用类型 - async： 异步 - sync：同步
	InvocationType ApiFuncInvocationType `json:"invocation_type"`

	// 函数版本   当函数别名URN和函数版本同时传入时，函数版本将被忽略，只会使用函数别名URN
	Version *string `json:"version,omitempty"`

	// 函数别名URN   当函数别名URN和函数版本同时传入时，函数版本将被忽略，只会使用函数别名URN
	AliasUrn *string `json:"alias_urn,omitempty"`

	// 服务集成请求后端服务的超时时间。最大超时时间可通过实例特性backend_timeout配置修改，可修改的上限为600000  单位：毫秒。
	Timeout int32 `json:"timeout"`

	// 后端自定义认证ID
	AuthorizerId *string `json:"authorizer_id,omitempty"`

	// 编号
	Id *string `json:"id,omitempty"`

	// 注册时间
	RegisterTime *sdktime.SdkTime `json:"register_time,omitempty"`

	// 后端状态   - 1： 有效
	Status *ApiFuncStatus `json:"status,omitempty"`

	// 修改时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`
}

func (o ApiFunc) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiFunc struct{}"
	}

	return strings.Join([]string{"ApiFunc", string(data)}, " ")
}

type ApiFuncNetworkType struct {
	value string
}

type ApiFuncNetworkTypeEnum struct {
	V1 ApiFuncNetworkType
	V2 ApiFuncNetworkType
}

func GetApiFuncNetworkTypeEnum() ApiFuncNetworkTypeEnum {
	return ApiFuncNetworkTypeEnum{
		V1: ApiFuncNetworkType{
			value: "V1",
		},
		V2: ApiFuncNetworkType{
			value: "V2",
		},
	}
}

func (c ApiFuncNetworkType) Value() string {
	return c.value
}

func (c ApiFuncNetworkType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiFuncNetworkType) UnmarshalJSON(b []byte) error {
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

type ApiFuncInvocationType struct {
	value string
}

type ApiFuncInvocationTypeEnum struct {
	ASYNC ApiFuncInvocationType
	SYNC  ApiFuncInvocationType
}

func GetApiFuncInvocationTypeEnum() ApiFuncInvocationTypeEnum {
	return ApiFuncInvocationTypeEnum{
		ASYNC: ApiFuncInvocationType{
			value: "async",
		},
		SYNC: ApiFuncInvocationType{
			value: "sync",
		},
	}
}

func (c ApiFuncInvocationType) Value() string {
	return c.value
}

func (c ApiFuncInvocationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiFuncInvocationType) UnmarshalJSON(b []byte) error {
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

type ApiFuncStatus struct {
	value int32
}

type ApiFuncStatusEnum struct {
	E_1 ApiFuncStatus
}

func GetApiFuncStatusEnum() ApiFuncStatusEnum {
	return ApiFuncStatusEnum{
		E_1: ApiFuncStatus{
			value: 1,
		},
	}
}

func (c ApiFuncStatus) Value() int32 {
	return c.value
}

func (c ApiFuncStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiFuncStatus) UnmarshalJSON(b []byte) error {
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
