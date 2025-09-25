package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AssetValueHostNumInfo 资产重要性的主机数量信息
type AssetValueHostNumInfo struct {

	// 重要性类型
	ValueType *AssetValueHostNumInfoValueType `json:"value_type,omitempty"`

	// 主机数量
	HostNum *int32 `json:"host_num,omitempty"`
}

func (o AssetValueHostNumInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssetValueHostNumInfo struct{}"
	}

	return strings.Join([]string{"AssetValueHostNumInfo", string(data)}, " ")
}

type AssetValueHostNumInfoValueType struct {
	value string
}

type AssetValueHostNumInfoValueTypeEnum struct {
	IMPORTANT AssetValueHostNumInfoValueType
	COMMON    AssetValueHostNumInfoValueType
	TEST      AssetValueHostNumInfoValueType
}

func GetAssetValueHostNumInfoValueTypeEnum() AssetValueHostNumInfoValueTypeEnum {
	return AssetValueHostNumInfoValueTypeEnum{
		IMPORTANT: AssetValueHostNumInfoValueType{
			value: "important",
		},
		COMMON: AssetValueHostNumInfoValueType{
			value: "common",
		},
		TEST: AssetValueHostNumInfoValueType{
			value: "test",
		},
	}
}

func (c AssetValueHostNumInfoValueType) Value() string {
	return c.value
}

func (c AssetValueHostNumInfoValueType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AssetValueHostNumInfoValueType) UnmarshalJSON(b []byte) error {
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
