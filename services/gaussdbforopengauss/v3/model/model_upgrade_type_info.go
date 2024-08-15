package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpgradeTypeInfo 升级类型信息。
type UpgradeTypeInfo struct {

	// 升级类型,grey=灰度升级,inplace=就地升级,hotfix=热补丁升级。
	UpgradeType *UpgradeTypeInfoUpgradeType `json:"upgrade_type,omitempty"`

	// 可用，不可用。
	Enable *bool `json:"enable,omitempty"`

	// 升级操作列表。
	UpgradeActionList *[]UpgradeActionInfo `json:"upgrade_action_list,omitempty"`

	// 是否正在进行AZ内并行升级。 -true：当前实例处于灰度升级的升级待观察升级方式中，已选择了AZ内并行升级方式，后续无法更改。 -false：当前实例处于升级流程中，未选择AZ内并行升级的方式，后续无法更改。null：当前实例尚未处于升级流程中。
	IsParallelUpgrade *bool `json:"is_parallel_upgrade,omitempty"`
}

func (o UpgradeTypeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeTypeInfo struct{}"
	}

	return strings.Join([]string{"UpgradeTypeInfo", string(data)}, " ")
}

type UpgradeTypeInfoUpgradeType struct {
	value string
}

type UpgradeTypeInfoUpgradeTypeEnum struct {
	GREY    UpgradeTypeInfoUpgradeType
	INPLACE UpgradeTypeInfoUpgradeType
	HOTFIX  UpgradeTypeInfoUpgradeType
}

func GetUpgradeTypeInfoUpgradeTypeEnum() UpgradeTypeInfoUpgradeTypeEnum {
	return UpgradeTypeInfoUpgradeTypeEnum{
		GREY: UpgradeTypeInfoUpgradeType{
			value: "grey",
		},
		INPLACE: UpgradeTypeInfoUpgradeType{
			value: "inplace",
		},
		HOTFIX: UpgradeTypeInfoUpgradeType{
			value: "hotfix",
		},
	}
}

func (c UpgradeTypeInfoUpgradeType) Value() string {
	return c.value
}

func (c UpgradeTypeInfoUpgradeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpgradeTypeInfoUpgradeType) UnmarshalJSON(b []byte) error {
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
