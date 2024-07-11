package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateHostClusterRequestBody 主机集群详细信息
type CreateHostClusterRequestBody struct {

	// 主机集群名
	Name string `json:"name"`

	// 项目id
	ProjectId string `json:"project_id"`

	// 操作系统：windows|linux
	Os CreateHostClusterRequestBodyOs `json:"os"`

	// slave集群id，默认为null时使用默认slave集群，用户自定义slave时为slave集群id
	SlaveClusterId *string `json:"slave_cluster_id,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 主机集群是否为代理机接入模式， 1：是 0：否
	IsProxyMode int32 `json:"is_proxy_mode"`
}

func (o CreateHostClusterRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHostClusterRequestBody struct{}"
	}

	return strings.Join([]string{"CreateHostClusterRequestBody", string(data)}, " ")
}

type CreateHostClusterRequestBodyOs struct {
	value string
}

type CreateHostClusterRequestBodyOsEnum struct {
	WINDOWS CreateHostClusterRequestBodyOs
	LINUX   CreateHostClusterRequestBodyOs
}

func GetCreateHostClusterRequestBodyOsEnum() CreateHostClusterRequestBodyOsEnum {
	return CreateHostClusterRequestBodyOsEnum{
		WINDOWS: CreateHostClusterRequestBodyOs{
			value: "windows",
		},
		LINUX: CreateHostClusterRequestBodyOs{
			value: "linux",
		},
	}
}

func (c CreateHostClusterRequestBodyOs) Value() string {
	return c.value
}

func (c CreateHostClusterRequestBodyOs) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateHostClusterRequestBodyOs) UnmarshalJSON(b []byte) error {
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
