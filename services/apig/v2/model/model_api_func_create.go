package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 函数后端详情
type ApiFuncCreate struct {

	// 函数URN
	FunctionUrn string `json:"function_urn" xml:"function_urn"`

	// 描述信息。长度不超过255个字符 > 中文字符必须为UTF-8或者unicode编码。
	Remark *string `json:"remark,omitempty" xml:"remark"`

	// 调用类型 - async： 异步 - sync：同步
	InvocationType ApiFuncCreateInvocationType `json:"invocation_type" xml:"invocation_type"`

	// 版本。
	Version *string `json:"version,omitempty" xml:"version"`

	// API网关请求后端服务的超时时间。最大超时时间可通过实例特性backend_timeout配置修改，可修改的上限为600000。  单位：毫秒。
	Timeout int32 `json:"timeout" xml:"timeout"`

	// 后端自定义认证ID
	AuthorizerId *string `json:"authorizer_id,omitempty" xml:"authorizer_id"`
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
