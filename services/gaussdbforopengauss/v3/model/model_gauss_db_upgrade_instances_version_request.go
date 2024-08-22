package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// GaussDbUpgradeInstancesVersionRequest GaussDB批量实例版本升级接口传参参数。
type GaussDbUpgradeInstancesVersionRequest struct {

	// 批量实例ID。
	InstanceIds []string `json:"instance_ids"`

	// 实例升级类型，包括就地升级，灰度升级、热补丁升级三种。
	UpgradeType GaussDbUpgradeInstancesVersionRequestUpgradeType `json:"upgrade_type"`

	// 实例升级操作，就地升级无需传值。灰度升级包括升级自动提交，升级待观察，提交升级，升级回退四种。
	UpgradeAction *GaussDbUpgradeInstancesVersionRequestUpgradeAction `json:"upgrade_action,omitempty"`

	// 批量实例升级目标版本，非必填。如果未传值灰度升级和就地升级默认升级到当前实例的优选版本，热补丁升级无需传该值，默认升级实例所有可升级热补丁。
	TargetVersion *string `json:"target_version,omitempty"`
}

func (o GaussDbUpgradeInstancesVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GaussDbUpgradeInstancesVersionRequest struct{}"
	}

	return strings.Join([]string{"GaussDbUpgradeInstancesVersionRequest", string(data)}, " ")
}

type GaussDbUpgradeInstancesVersionRequestUpgradeType struct {
	value string
}

type GaussDbUpgradeInstancesVersionRequestUpgradeTypeEnum struct {
	INPLACE GaussDbUpgradeInstancesVersionRequestUpgradeType
	GREY    GaussDbUpgradeInstancesVersionRequestUpgradeType
	HOTFIX  GaussDbUpgradeInstancesVersionRequestUpgradeType
}

func GetGaussDbUpgradeInstancesVersionRequestUpgradeTypeEnum() GaussDbUpgradeInstancesVersionRequestUpgradeTypeEnum {
	return GaussDbUpgradeInstancesVersionRequestUpgradeTypeEnum{
		INPLACE: GaussDbUpgradeInstancesVersionRequestUpgradeType{
			value: "inplace",
		},
		GREY: GaussDbUpgradeInstancesVersionRequestUpgradeType{
			value: "grey",
		},
		HOTFIX: GaussDbUpgradeInstancesVersionRequestUpgradeType{
			value: "hotfix",
		},
	}
}

func (c GaussDbUpgradeInstancesVersionRequestUpgradeType) Value() string {
	return c.value
}

func (c GaussDbUpgradeInstancesVersionRequestUpgradeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GaussDbUpgradeInstancesVersionRequestUpgradeType) UnmarshalJSON(b []byte) error {
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

type GaussDbUpgradeInstancesVersionRequestUpgradeAction struct {
	value string
}

type GaussDbUpgradeInstancesVersionRequestUpgradeActionEnum struct {
	UPGRADE_AUTO_COMMIT GaussDbUpgradeInstancesVersionRequestUpgradeAction
	UPGRADE             GaussDbUpgradeInstancesVersionRequestUpgradeAction
	COMMIT              GaussDbUpgradeInstancesVersionRequestUpgradeAction
	ROLLBACK            GaussDbUpgradeInstancesVersionRequestUpgradeAction
}

func GetGaussDbUpgradeInstancesVersionRequestUpgradeActionEnum() GaussDbUpgradeInstancesVersionRequestUpgradeActionEnum {
	return GaussDbUpgradeInstancesVersionRequestUpgradeActionEnum{
		UPGRADE_AUTO_COMMIT: GaussDbUpgradeInstancesVersionRequestUpgradeAction{
			value: "upgradeAutoCommit",
		},
		UPGRADE: GaussDbUpgradeInstancesVersionRequestUpgradeAction{
			value: "upgrade",
		},
		COMMIT: GaussDbUpgradeInstancesVersionRequestUpgradeAction{
			value: "commit",
		},
		ROLLBACK: GaussDbUpgradeInstancesVersionRequestUpgradeAction{
			value: "rollback",
		},
	}
}

func (c GaussDbUpgradeInstancesVersionRequestUpgradeAction) Value() string {
	return c.value
}

func (c GaussDbUpgradeInstancesVersionRequestUpgradeAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GaussDbUpgradeInstancesVersionRequestUpgradeAction) UnmarshalJSON(b []byte) error {
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
