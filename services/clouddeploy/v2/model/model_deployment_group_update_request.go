package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

//
type DeploymentGroupUpdateRequest struct {
	// 主机组名

	Name string `json:"name"`
	// 描述

	Description *string `json:"description,omitempty"`
	// 自定义slave资源池id

	SlaveClusterId *string `json:"slave_cluster_id,omitempty"`
	// 自动连通性验证 0不执行 1每日 2每周

	AutoConnectionTestSwitch *DeploymentGroupUpdateRequestAutoConnectionTestSwitch `json:"auto_connection_test_switch,omitempty"`
}

func (o DeploymentGroupUpdateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeploymentGroupUpdateRequest struct{}"
	}

	return strings.Join([]string{"DeploymentGroupUpdateRequest", string(data)}, " ")
}

type DeploymentGroupUpdateRequestAutoConnectionTestSwitch struct {
	value int32
}

type DeploymentGroupUpdateRequestAutoConnectionTestSwitchEnum struct {
	E_0 DeploymentGroupUpdateRequestAutoConnectionTestSwitch
	E_1 DeploymentGroupUpdateRequestAutoConnectionTestSwitch
	E_2 DeploymentGroupUpdateRequestAutoConnectionTestSwitch
}

func GetDeploymentGroupUpdateRequestAutoConnectionTestSwitchEnum() DeploymentGroupUpdateRequestAutoConnectionTestSwitchEnum {
	return DeploymentGroupUpdateRequestAutoConnectionTestSwitchEnum{
		E_0: DeploymentGroupUpdateRequestAutoConnectionTestSwitch{
			value: 0,
		}, E_1: DeploymentGroupUpdateRequestAutoConnectionTestSwitch{
			value: 1,
		}, E_2: DeploymentGroupUpdateRequestAutoConnectionTestSwitch{
			value: 2,
		},
	}
}

func (c DeploymentGroupUpdateRequestAutoConnectionTestSwitch) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeploymentGroupUpdateRequestAutoConnectionTestSwitch) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
