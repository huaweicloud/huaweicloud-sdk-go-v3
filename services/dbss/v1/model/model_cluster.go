package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Cluster struct {
	ActivateInfo *InstanceActivateInfo `json:"activate_info,omitempty"`

	// 计费模式  - Period：包周期  - Demand：按需
	ChargeModel *ClusterChargeModel `json:"charge_model,omitempty"`

	// 备注信息
	Comment *string `json:"comment,omitempty"`

	// 创建时间
	Created *int64 `json:"created,omitempty"`

	// 部署方式  - CLOUD：云上  - OUTSIDE：云外
	DeployMode *ClusterDeployMode `json:"deploy_mode,omitempty"`

	Detail *ServerList `json:"detail,omitempty"`

	// 企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 过期时间
	Expired *int64 `json:"expired,omitempty"`

	// 浮动IP
	FloatIp *string `json:"float_ip,omitempty"`

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 是否激活备用
	IsActiveStandby *bool `json:"is_active_standby,omitempty"`

	// 使用天数
	KeepDays *string `json:"keep_days,omitempty"`

	// 实例名称
	Name *string `json:"name,omitempty"`

	// 最新版本
	NewVersion *string `json:"new_version,omitempty"`

	// 公网IP
	PublicIp *string `json:"public_ip,omitempty"`

	// 剩余天数
	RemainDays *string `json:"remain_days,omitempty"`
}

func (o Cluster) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Cluster struct{}"
	}

	return strings.Join([]string{"Cluster", string(data)}, " ")
}

type ClusterChargeModel struct {
	value string
}

type ClusterChargeModelEnum struct {
	PERIOD ClusterChargeModel
	DEMAND ClusterChargeModel
	TEST   ClusterChargeModel
}

func GetClusterChargeModelEnum() ClusterChargeModelEnum {
	return ClusterChargeModelEnum{
		PERIOD: ClusterChargeModel{
			value: "Period",
		},
		DEMAND: ClusterChargeModel{
			value: "Demand",
		},
		TEST: ClusterChargeModel{
			value: "Test",
		},
	}
}

func (c ClusterChargeModel) Value() string {
	return c.value
}

func (c ClusterChargeModel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ClusterChargeModel) UnmarshalJSON(b []byte) error {
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

type ClusterDeployMode struct {
	value string
}

type ClusterDeployModeEnum struct {
	CLOUD   ClusterDeployMode
	OUTSIDE ClusterDeployMode
}

func GetClusterDeployModeEnum() ClusterDeployModeEnum {
	return ClusterDeployModeEnum{
		CLOUD: ClusterDeployMode{
			value: "CLOUD",
		},
		OUTSIDE: ClusterDeployMode{
			value: "OUTSIDE",
		},
	}
}

func (c ClusterDeployMode) Value() string {
	return c.value
}

func (c ClusterDeployMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ClusterDeployMode) UnmarshalJSON(b []byte) error {
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
