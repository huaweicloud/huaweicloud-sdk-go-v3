package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type DimensionInfo struct {

	// 维度名称，多维度用逗号分隔，各服务支持的维度可参考：“[服务维度名称](https://support.huaweicloud.com/usermanual-ces/zh-cn_topic_0202622212.html)”
	Name string `json:"name"`

	// 资源类型, all_instances: 全部资源, specific_instances: 指定资源
	FilterType DimensionInfoFilterType `json:"filter_type"`

	// 维度值列表
	Values *[]string `json:"values,omitempty"`
}

func (o DimensionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DimensionInfo struct{}"
	}

	return strings.Join([]string{"DimensionInfo", string(data)}, " ")
}

type DimensionInfoFilterType struct {
	value string
}

type DimensionInfoFilterTypeEnum struct {
	ALL_INSTANCES      DimensionInfoFilterType
	SPECIFIC_INSTANCES DimensionInfoFilterType
}

func GetDimensionInfoFilterTypeEnum() DimensionInfoFilterTypeEnum {
	return DimensionInfoFilterTypeEnum{
		ALL_INSTANCES: DimensionInfoFilterType{
			value: "all_instances",
		},
		SPECIFIC_INSTANCES: DimensionInfoFilterType{
			value: "specific_instances",
		},
	}
}

func (c DimensionInfoFilterType) Value() string {
	return c.value
}

func (c DimensionInfoFilterType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DimensionInfoFilterType) UnmarshalJSON(b []byte) error {
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
