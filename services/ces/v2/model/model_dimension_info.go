package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DimensionInfo **参数解释** 维度信息列表。       **约束限制** 不涉及
type DimensionInfo struct {

	// **参数解释**： 维度名称。 **约束限制** 不涉及 **取值范围**： 多维度用逗号分隔，各服务资源的维度名称，请参阅具体云服务的文档。您可以直接从[[支持监控的服务列表](https://support.huaweicloud.com/api-ces/ces_03_0059.html)](tag:hc)[[支持监控的服务列表](https://support.huaweicloud.com/intl/en-us/api-ces/ces_03_0059.html)](tag:hk)[[支持监控的服务列表](https://support.huaweicloud.com/eu/en-us/api-ces/ces_03_0059.html)](tag:hws_eu)[[支持监控的服务列表](ces_03_0059.xml)](tag:ax,cmcc,ctc,dt,dt_test,hcso_dt,fcs,fcs_vm,mix,g42,hk_g42,hk_sbc,hk_tm,hk_vdf,hws_ocb,ocb,sbc,srg)页面导航至相应文档。。必须以字母开头，只能包含0-9/a-z/A-Z/_/-，多维度用\",\"分隔，每个维度的最大长度为32。总长度为[1,131]个字符。目前最大支持4个维度。举例：单维度场景：instance_id；多维度场景：instance_id,disk **默认取值** 不涉及
	Name string `json:"name"`

	// **参数解释**： 资源类型。 **约束限制** view参数取值为柱状图 条形图 环形柱状图 表格时，则filter_type参数不支持指定资源 **取值范围**： - all_instances: 全部资源 - specific_instances: 指定资源 **默认取值** 不涉及
	FilterType DimensionInfoFilterType `json:"filter_type"`

	// **参数解释**： 维度值列表。 **约束限制** 包含的维度值对象个数为[0,200]
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
