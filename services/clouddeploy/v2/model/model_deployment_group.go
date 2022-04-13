package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 主机组详细信息
type DeploymentGroup struct {
	// 主机组名

	Name string `json:"name"`
	// 局点信息

	RegionName string `json:"region_name"`
	// devcloud项目id

	ProjectId string `json:"project_id"`
	// 操作信息：windows|linux

	Os DeploymentGroupOs `json:"os"`
	// slave集群id，默认为null时使用devcloud八爪鱼slave集群，用户自定义slave时为slave集群id

	SlaveClusterId *string `json:"slave_cluster_id,omitempty"`
	// 描述

	Description *string `json:"description,omitempty"`
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

func (c DeploymentGroupOs) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeploymentGroupOs) UnmarshalJSON(b []byte) error {
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
