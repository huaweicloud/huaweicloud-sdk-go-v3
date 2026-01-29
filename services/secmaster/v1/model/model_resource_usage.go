package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ResourceUsage 资源用量，当接口请求参数usage=true时会返回具体的使用量信息，其它场景不返回
type ResourceUsage struct {

	// 使用量单位，OPS 次，MB 流量体积MB，GB 流量体积GB
	Unit *ResourceUsageUnit `json:"unit,omitempty"`

	// 资源类型名称
	ResourceTypeName *string `json:"resource_type_name,omitempty"`

	// 源资源规格编码
	SourceResourceSpecCode *string `json:"source_resource_spec_code,omitempty"`

	// 源资源规格编码
	ResourceSpecCode *interface{} `json:"resource_spec_code,omitempty"`

	// 源资源类型编码
	SourceType *string `json:"source_type,omitempty"`

	// 用量百分比
	UsedPercent *float64 `json:"used_percent,omitempty"`

	// 配额总量
	Quota *float64 `json:"quota,omitempty"`

	// 已用量
	Used *float64 `json:"used,omitempty"`

	// 剩余量
	Free *float64 `json:"free,omitempty"`
}

func (o ResourceUsage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceUsage struct{}"
	}

	return strings.Join([]string{"ResourceUsage", string(data)}, " ")
}

type ResourceUsageUnit struct {
	value string
}

type ResourceUsageUnitEnum struct {
	OPS ResourceUsageUnit
	MB  ResourceUsageUnit
	GB  ResourceUsageUnit
}

func GetResourceUsageUnitEnum() ResourceUsageUnitEnum {
	return ResourceUsageUnitEnum{
		OPS: ResourceUsageUnit{
			value: "OPS",
		},
		MB: ResourceUsageUnit{
			value: "MB",
		},
		GB: ResourceUsageUnit{
			value: "GB",
		},
	}
}

func (c ResourceUsageUnit) Value() string {
	return c.value
}

func (c ResourceUsageUnit) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResourceUsageUnit) UnmarshalJSON(b []byte) error {
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
