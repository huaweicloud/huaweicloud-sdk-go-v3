package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ExportInstancesTaskBody struct {

	// 导出实例列表，如果为空，则会导出满足其余参数条件的所有实例
	IncludedInstances *[]string `json:"included_instances,omitempty"`

	// 局点名称，用于导出的文件名命名
	Region string `json:"region"`

	// 按照实例名称筛选实例
	Name *string `json:"name,omitempty"`

	// 按照实例规格筛选实例
	Capacity *string `json:"capacity,omitempty"`

	// 按照实例ID筛选实例
	InstanceId *string `json:"instance_id,omitempty"`

	// 按照ip筛选实例
	Ip *string `json:"ip,omitempty"`

	// 按照可用区筛选实例
	AvailableZone *string `json:"available_zone,omitempty"`

	// 按照实例状态筛选实例
	Status *string `json:"status,omitempty"`

	// 按照产品类型筛选实例，generic-普通版本，enterprise-企业版
	ProductType *ExportInstancesTaskBodyProductType `json:"product_type,omitempty"`

	// 按照实例类型筛选实例
	CacheMode *string `json:"cache_mode,omitempty"`

	// 按照缓存引擎筛选实例
	Engine *string `json:"engine,omitempty"`

	// 按照缓存引擎版本筛选实例
	EngineVersion *string `json:"engine_version,omitempty"`

	// 按照CPU类型筛选实例
	CpuType *string `json:"cpu_type,omitempty"`

	// 按照企业项目ID筛选实例
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 按照计费方式筛选实例
	ChargingMode *string `json:"charging_mode,omitempty"`

	// 按照标签筛选实例
	Tags *string `json:"tags,omitempty"`
}

func (o ExportInstancesTaskBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportInstancesTaskBody struct{}"
	}

	return strings.Join([]string{"ExportInstancesTaskBody", string(data)}, " ")
}

type ExportInstancesTaskBodyProductType struct {
	value string
}

type ExportInstancesTaskBodyProductTypeEnum struct {
	GENERIC    ExportInstancesTaskBodyProductType
	ENTERPRISE ExportInstancesTaskBodyProductType
}

func GetExportInstancesTaskBodyProductTypeEnum() ExportInstancesTaskBodyProductTypeEnum {
	return ExportInstancesTaskBodyProductTypeEnum{
		GENERIC: ExportInstancesTaskBodyProductType{
			value: "generic",
		},
		ENTERPRISE: ExportInstancesTaskBodyProductType{
			value: "enterprise",
		},
	}
}

func (c ExportInstancesTaskBodyProductType) Value() string {
	return c.value
}

func (c ExportInstancesTaskBodyProductType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportInstancesTaskBodyProductType) UnmarshalJSON(b []byte) error {
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
