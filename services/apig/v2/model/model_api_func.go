package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"errors"
	"strings"
)

// 函数工作流后端详情
type ApiFunc struct {
	// 函数URN

	FunctionUrn string `json:"function_urn"`
	// 描述信息。长度不超过255个字符 > 中文字符必须为UTF-8或者unicode编码。

	Remark *string `json:"remark,omitempty"`
	// 调用类型 - async： 异步 - sync：同步

	InvocationType ApiFuncInvocationType `json:"invocation_type"`
	// 版本。

	Version *string `json:"version,omitempty"`
	// API网关请求后端服务的超时时间。最大超时时间可通过实例特性backend_timeout配置修改，可修改的上限为600000。  单位：毫秒。

	Timeout int32 `json:"timeout"`
	// 后端自定义认证ID

	AuthorizerId *string `json:"authorizer_id,omitempty"`
	// 编号

	Id *string `json:"id,omitempty"`
	// 注册时间

	RegisterTime *sdktime.SdkTime `json:"register_time,omitempty"`
	// 后端状态   - 1： 有效

	Status *int32 `json:"status,omitempty"`
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

func (c ApiFuncInvocationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiFuncInvocationType) UnmarshalJSON(b []byte) error {
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
