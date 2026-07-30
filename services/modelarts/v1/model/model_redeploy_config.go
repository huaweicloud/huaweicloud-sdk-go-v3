package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type RedeployConfig struct {

	// 节点的重部署类型。若节点状态为不可用，将无法进行SOFT模式，只能进行HARD模式，HARD模式包含节点重置操作，会导致本地盘及云盘上的全部数据丢失，请谨慎操作
	Type *RedeployConfigType `json:"type,omitempty"`

	// 静默修复开关。开启autoFlow开关时，如重部署失败系统将自动流转至\"系统维护\"或发起\"二次重部署\"，并产生新的计划事件，该过程自动授权，无需二次授权
	AutoFlow *RedeployConfigAutoFlow `json:"autoFlow,omitempty"`
}

func (o RedeployConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RedeployConfig struct{}"
	}

	return strings.Join([]string{"RedeployConfig", string(data)}, " ")
}

type RedeployConfigType struct {
	value string
}

type RedeployConfigTypeEnum struct {
	SOFT RedeployConfigType
	HARD RedeployConfigType
}

func GetRedeployConfigTypeEnum() RedeployConfigTypeEnum {
	return RedeployConfigTypeEnum{
		SOFT: RedeployConfigType{
			value: "SOFT",
		},
		HARD: RedeployConfigType{
			value: "HARD",
		},
	}
}

func (c RedeployConfigType) Value() string {
	return c.value
}

func (c RedeployConfigType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RedeployConfigType) UnmarshalJSON(b []byte) error {
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

type RedeployConfigAutoFlow struct {
	value string
}

type RedeployConfigAutoFlowEnum struct {
	TRUE  RedeployConfigAutoFlow
	FALSE RedeployConfigAutoFlow
}

func GetRedeployConfigAutoFlowEnum() RedeployConfigAutoFlowEnum {
	return RedeployConfigAutoFlowEnum{
		TRUE: RedeployConfigAutoFlow{
			value: "true",
		},
		FALSE: RedeployConfigAutoFlow{
			value: "false",
		},
	}
}

func (c RedeployConfigAutoFlow) Value() string {
	return c.value
}

func (c RedeployConfigAutoFlow) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RedeployConfigAutoFlow) UnmarshalJSON(b []byte) error {
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
