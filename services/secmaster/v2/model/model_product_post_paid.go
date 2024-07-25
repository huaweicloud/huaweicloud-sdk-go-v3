package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ProductPostPaid struct {

	// ID标识，同一次询价中不能重复，用于标识返回询价结果和请求的映射关系
	Id string `json:"id"`

	// 产品Id，通过向CBC询价获取该商品的标识
	ProductId string `json:"product_id"`

	// 云服务类型，固定值为hws.service.type.sa
	CloudServiceType string `json:"cloud_service_type"`

	// 用户购买云服务产品的资源类型，例如SecMaster中的典型场景配置，资源类型为hws.resource.type.secmaster.typical
	ResourceType string `json:"resource_type"`

	// 用户购买云服务产品的资源规格，例如安全云脑中的的基础版，资源规格为secmaster.basic
	ResourceSpecCode string `json:"resource_spec_code"`

	// 使用量单位标识，按需询价必填，例如按小时询价，使用量值为1，使用量单位为小时，枚举值如下： 4：小时 10：GB（带宽按流量询价使用） 11：MB（带宽按流量询价使用）
	UsageMeasureId ProductPostPaidUsageMeasureId `json:"usage_measure_id"`

	// 使用量值，按需询价必填，例如按小时询价，使用量值为1，使用量单位为小时
	UsageValue float32 `json:"usage_value"`

	// 配额个数
	ResourceSize int32 `json:"resource_size"`

	// 使用量因子，按需计费必填，取值和话单中的使用量因子一致，云服务和使用量因子对应关系如下： 典型场景配置：Duration 态势管理：duration 安全编排：count 智能分析：flow
	UsageFactor string `json:"usage_factor"`

	// 资源id，仅在增加配额的时候传入
	ResourceId *string `json:"resource_id,omitempty"`
}

func (o ProductPostPaid) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProductPostPaid struct{}"
	}

	return strings.Join([]string{"ProductPostPaid", string(data)}, " ")
}

type ProductPostPaidUsageMeasureId struct {
	value int32
}

type ProductPostPaidUsageMeasureIdEnum struct {
	E_4  ProductPostPaidUsageMeasureId
	E_10 ProductPostPaidUsageMeasureId
	E_11 ProductPostPaidUsageMeasureId
}

func GetProductPostPaidUsageMeasureIdEnum() ProductPostPaidUsageMeasureIdEnum {
	return ProductPostPaidUsageMeasureIdEnum{
		E_4: ProductPostPaidUsageMeasureId{
			value: 4,
		}, E_10: ProductPostPaidUsageMeasureId{
			value: 10,
		}, E_11: ProductPostPaidUsageMeasureId{
			value: 11,
		},
	}
}

func (c ProductPostPaidUsageMeasureId) Value() int32 {
	return c.value
}

func (c ProductPostPaidUsageMeasureId) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ProductPostPaidUsageMeasureId) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
