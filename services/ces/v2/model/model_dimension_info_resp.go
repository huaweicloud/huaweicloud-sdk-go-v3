package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DimensionInfoResp **参数描述** 维度信息列表。
type DimensionInfoResp struct {

	// **参数描述**： 维度名称。 **取值范围**： 多维度用逗号分隔，各服务支持的维度可参考：“[服务维度名称](ces_03_0059.xml)”。必须以字母开头，只能包含0-9/a-z/A-Z/_/-，多维度用\",\"分隔，每个维度的最大长度为32。总长度为[1,131]个字符。目前最大支持4个维度。举例：单维度场景：instance_id；多维度场景：instance_id,disk
	Name *string `json:"name,omitempty"`

	// **参数描述**： 资源类型。 **取值范围**： - all_instances: 全部资源 - specific_instances: 指定资源
	FilterType *DimensionInfoRespFilterType `json:"filter_type,omitempty"`

	// **参数描述**： 维度值列表。
	Values *[]string `json:"values,omitempty"`
}

func (o DimensionInfoResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DimensionInfoResp struct{}"
	}

	return strings.Join([]string{"DimensionInfoResp", string(data)}, " ")
}

type DimensionInfoRespFilterType struct {
	value string
}

type DimensionInfoRespFilterTypeEnum struct {
	ALL_INSTANCES      DimensionInfoRespFilterType
	SPECIFIC_INSTANCES DimensionInfoRespFilterType
}

func GetDimensionInfoRespFilterTypeEnum() DimensionInfoRespFilterTypeEnum {
	return DimensionInfoRespFilterTypeEnum{
		ALL_INSTANCES: DimensionInfoRespFilterType{
			value: "all_instances",
		},
		SPECIFIC_INSTANCES: DimensionInfoRespFilterType{
			value: "specific_instances",
		},
	}
}

func (c DimensionInfoRespFilterType) Value() string {
	return c.value
}

func (c DimensionInfoRespFilterType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DimensionInfoRespFilterType) UnmarshalJSON(b []byte) error {
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
