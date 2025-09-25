package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AssetResourceStatusEnum **参数解释**： 计费资产状态。 **约束限制**： 不涉及 **取值范围**： - Normal：正常 - Freeze：冻结 - Deleted：删除 **默认取值**： 不涉及
type AssetResourceStatusEnum struct {
	value string
}

type AssetResourceStatusEnumEnum struct {
	NORMAL  AssetResourceStatusEnum
	FREEZE  AssetResourceStatusEnum
	DELETED AssetResourceStatusEnum
}

func GetAssetResourceStatusEnumEnum() AssetResourceStatusEnumEnum {
	return AssetResourceStatusEnumEnum{
		NORMAL: AssetResourceStatusEnum{
			value: "Normal",
		},
		FREEZE: AssetResourceStatusEnum{
			value: "Freeze",
		},
		DELETED: AssetResourceStatusEnum{
			value: "Deleted",
		},
	}
}

func (c AssetResourceStatusEnum) Value() string {
	return c.value
}

func (c AssetResourceStatusEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AssetResourceStatusEnum) UnmarshalJSON(b []byte) error {
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
