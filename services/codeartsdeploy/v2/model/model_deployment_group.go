package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeploymentGroup 主机集群详细信息
type DeploymentGroup struct {

	// 主机集群名
	Name string `json:"name"`

	// 局点信息
	RegionName string `json:"region_name"`

	// 项目id
	ProjectId string `json:"project_id"`

	// 操作信息：windows|linux
	Os DeploymentGroupOs `json:"os"`

	// slave集群id，默认为null时使用默认slave集群，用户自定义slave时为slave集群id
	SlaveClusterId *string `json:"slave_cluster_id,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 主机集群是否为代理类型
	IsProxyMode *int32 `json:"is_proxy_mode,omitempty"`
}

func (o DeploymentGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeploymentGroup struct{}"
	}

	return strings.Join([]string{"DeploymentGroup", string(data)}, " ")
}

type DeploymentGroupOs struct {
	value string
}

type DeploymentGroupOsEnum struct {
	WINDOWS DeploymentGroupOs
	LINUX   DeploymentGroupOs
}

func GetDeploymentGroupOsEnum() DeploymentGroupOsEnum {
	return DeploymentGroupOsEnum{
		WINDOWS: DeploymentGroupOs{
			value: "windows",
		},
		LINUX: DeploymentGroupOs{
			value: "linux",
		},
	}
}

func (c DeploymentGroupOs) Value() string {
	return c.value
}

func (c DeploymentGroupOs) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeploymentGroupOs) UnmarshalJSON(b []byte) error {
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
