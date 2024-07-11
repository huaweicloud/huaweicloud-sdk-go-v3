package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// EnvironmentRequestBody 新建和编辑环境的请求体
type EnvironmentRequestBody struct {

	// 项目id
	ProjectId string `json:"project_id"`

	// 环境名称
	Name string `json:"name"`

	// 部署类型：0表示主机, 1表示kubernetes
	DeployType int32 `json:"deploy_type"`

	// 操作系统：windows|linux，需要和主机集群保持一致
	Os EnvironmentRequestBodyOs `json:"os"`

	// 环境描述
	Description *string `json:"description,omitempty"`
}

func (o EnvironmentRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnvironmentRequestBody struct{}"
	}

	return strings.Join([]string{"EnvironmentRequestBody", string(data)}, " ")
}

type EnvironmentRequestBodyOs struct {
	value string
}

type EnvironmentRequestBodyOsEnum struct {
	WINDOWS EnvironmentRequestBodyOs
	LINUX   EnvironmentRequestBodyOs
}

func GetEnvironmentRequestBodyOsEnum() EnvironmentRequestBodyOsEnum {
	return EnvironmentRequestBodyOsEnum{
		WINDOWS: EnvironmentRequestBodyOs{
			value: "windows",
		},
		LINUX: EnvironmentRequestBodyOs{
			value: "linux",
		},
	}
}

func (c EnvironmentRequestBodyOs) Value() string {
	return c.value
}

func (c EnvironmentRequestBodyOs) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EnvironmentRequestBodyOs) UnmarshalJSON(b []byte) error {
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
