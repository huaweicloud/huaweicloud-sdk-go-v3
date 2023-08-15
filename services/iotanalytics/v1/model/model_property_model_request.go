package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PropertyModelRequest struct {

	// 属性名称，正则：\"^[a-zA-Z0-9_]{1,64}$\"
	Name string `json:"name"`

	// 属性显示名称，正则：\"^[\\\\u4E00-\\\\u9FA5A-Za-z0-9_@#.-]{1,64}$\"
	DisplayName *string `json:"display_name,omitempty"`

	// 属性类别，静态配置（static）、测量数据（measurement）、分析任务（analysis）
	SourceType string `json:"source_type"`

	DataSchema *DataSchema `json:"data_schema"`

	// 单位
	Unit *string `json:"unit,omitempty"`

	// 静态属性的值，如：1 1.1 \"value\" {\"name\":\"value\"}
	Value *interface{} `json:"value,omitempty"`

	// 属性是否为标签。资产ID、标签属性、时间戳三者组成属性数据唯一键，两条唯一键相同的属性数据以覆盖方式存储；一个模型中只能配置三个属性为标签，标签配置后标签不能删除，配置标签的属性也不能删除；只有integer、double、string类型的属性可以被配置为标签。示例： 资产ID asset1上依次上报如下六组数据： 资产ID 属性A（标签） 属性B    属性C 时间戳 asset1 valueA_1     valueB_1  valueC_1 T1 asset1 valueA_1     valueB_2  valueC_2 T2 asset1 valueA_2     valueB_3  valueC_3 T2 asset1 valueA_2     valueB_4  valueC_4 T2 asset1              valueB_5  valueC_5 T3 asset1              valueB_6  valueC_6 T3 根据唯一键规则最终存储为如下四组数据： 资产ID 属性A（标签） 属性B    属性C 时间戳 asset1 valueA_1     valueB_1  valueC_1 T1 asset1 valueA_1     valueB_2  valueC_2 T2 asset1 valueA_2     valueB_4  valueC_4 T2 asset1              valueB_6  valueC_6 T3
	IsTag *bool `json:"is_tag,omitempty"`
}

func (o PropertyModelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PropertyModelRequest struct{}"
	}

	return strings.Join([]string{"PropertyModelRequest", string(data)}, " ")
}
