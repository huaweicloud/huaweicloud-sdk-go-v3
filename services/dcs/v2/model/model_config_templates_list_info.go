package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ConfigTemplatesListInfo struct {

	// 模板ID
	TemplateId *string `json:"template_id,omitempty"`

	// 缓存实例类型。取值范围如下： - single：表示单机实例 - ha：表示主备实例 - cluster：表示cluster集群实例 - proxy：表示Proxy集群实例 [- ha_rw_split： 表示读写分离实例](tag:hws)
	CacheMode *string `json:"cache_mode,omitempty"`

	// 模板的描述信息
	Description *string `json:"description,omitempty"`

	// 缓存引擎：Redis[和Memcached](tag:hws,hws_hk,ocb,sbc,tm,ctc,cmcc)。
	Engine *string `json:"engine,omitempty"`

	// 缓存版本。  当缓存引擎为Redis时，取值为4.0或5.0。  [当缓存引擎为Memcached时，该字段为可选，取值为空。](tag:hws,hws_hk,ocb,sbc,tm,ctc,cmcc)
	EngineVersion *string `json:"engine_version,omitempty"`

	// 模板名称
	Name *string `json:"name,omitempty"`

	// 产品类型。 取值范围如下： - generic：普通版 - enterprise：企业版
	ProductType *ConfigTemplatesListInfoProductType `json:"product_type,omitempty"`

	// 存储类型。 取值范围如下： - DRAM - SSD
	StorageType *ConfigTemplatesListInfoStorageType `json:"storage_type,omitempty"`

	// 模板类型
	Type *string `json:"type,omitempty"`

	// 模板创建时间，仅在自定义参数模板中有意义，格式例如：2023-05-10T11:09:35.802Z
	CreatedAt *string `json:"created_at,omitempty"`
}

func (o ConfigTemplatesListInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigTemplatesListInfo struct{}"
	}

	return strings.Join([]string{"ConfigTemplatesListInfo", string(data)}, " ")
}

type ConfigTemplatesListInfoProductType struct {
	value string
}

type ConfigTemplatesListInfoProductTypeEnum struct {
	GENERIC    ConfigTemplatesListInfoProductType
	ENTERPRISE ConfigTemplatesListInfoProductType
}

func GetConfigTemplatesListInfoProductTypeEnum() ConfigTemplatesListInfoProductTypeEnum {
	return ConfigTemplatesListInfoProductTypeEnum{
		GENERIC: ConfigTemplatesListInfoProductType{
			value: "generic",
		},
		ENTERPRISE: ConfigTemplatesListInfoProductType{
			value: "enterprise",
		},
	}
}

func (c ConfigTemplatesListInfoProductType) Value() string {
	return c.value
}

func (c ConfigTemplatesListInfoProductType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConfigTemplatesListInfoProductType) UnmarshalJSON(b []byte) error {
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

type ConfigTemplatesListInfoStorageType struct {
	value string
}

type ConfigTemplatesListInfoStorageTypeEnum struct {
	DRAM ConfigTemplatesListInfoStorageType
	SSD  ConfigTemplatesListInfoStorageType
}

func GetConfigTemplatesListInfoStorageTypeEnum() ConfigTemplatesListInfoStorageTypeEnum {
	return ConfigTemplatesListInfoStorageTypeEnum{
		DRAM: ConfigTemplatesListInfoStorageType{
			value: "DRAM",
		},
		SSD: ConfigTemplatesListInfoStorageType{
			value: "SSD",
		},
	}
}

func (c ConfigTemplatesListInfoStorageType) Value() string {
	return c.value
}

func (c ConfigTemplatesListInfoStorageType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConfigTemplatesListInfoStorageType) UnmarshalJSON(b []byte) error {
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
