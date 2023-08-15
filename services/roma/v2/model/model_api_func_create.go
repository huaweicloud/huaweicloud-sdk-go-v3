package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ApiFuncCreate [函数工作流后端详情](tag:hws;hws_hk;hcs;fcs;g42;)[暂不支持](tag:Site)
type ApiFuncCreate struct {

	// 函数URN
	FunctionUrn string `json:"function_urn"`

	// 描述信息。 > 中文字符必须为UTF-8或者unicode编码。
	Remark *string `json:"remark,omitempty"`

	// 调用类型 - async： 异步 - sync：同步
	InvocationType ApiFuncCreateInvocationType `json:"invocation_type"`

	// 版本。
	Version *string `json:"version,omitempty"`

	// 函数别名URN  当函数别名URN和函数版本同时传入时，函数版本将被忽略，只会使用函数别名URN
	AliasUrn *string `json:"alias_urn,omitempty"`

	// ROMA Connect APIC请求后端服务的超时时间。最大超时时间可通过实例特性backend_timeout配置修改，可修改的上限为600000  单位：毫秒。
	Timeout int32 `json:"timeout"`

	// 后端自定义认证ID
	AuthorizerId *string `json:"authorizer_id,omitempty"`
}

func (o ApiFuncCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiFuncCreate struct{}"
	}

	return strings.Join([]string{"ApiFuncCreate", string(data)}, " ")
}

type ApiFuncCreateInvocationType struct {
	value string
}

type ApiFuncCreateInvocationTypeEnum struct {
	ASYNC ApiFuncCreateInvocationType
	SYNC  ApiFuncCreateInvocationType
}

func GetApiFuncCreateInvocationTypeEnum() ApiFuncCreateInvocationTypeEnum {
	return ApiFuncCreateInvocationTypeEnum{
		ASYNC: ApiFuncCreateInvocationType{
			value: "async",
		},
		SYNC: ApiFuncCreateInvocationType{
			value: "sync",
		},
	}
}

func (c ApiFuncCreateInvocationType) Value() string {
	return c.value
}

func (c ApiFuncCreateInvocationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiFuncCreateInvocationType) UnmarshalJSON(b []byte) error {
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
