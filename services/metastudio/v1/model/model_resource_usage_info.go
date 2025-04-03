package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ResourceUsageInfo 资源使用信息
type ResourceUsageInfo struct {

	// 资源类型。详见[资源类型](metastudio_02_0042.xml)。
	ResourceType *string `json:"resource_type,omitempty"`

	// 业务类型。 * VOICE_CLONE：声音制作 * SYNTHETICS_SOUND：声音合成 * ASSET_MANAGER：资产管理 * MODELING_2D：形象制作 * LIVE_2D：分身数字人视频直播 * VIDEO_2D：分身数字人视频制作 * CHAT_2D：分身数字人智能交互 * BUSINESS_CARD_2D：分身数字人名片 * PICTURE_2D：照片数字人视频 * MODELING_3D：3D照片建模 * VDS_3D：3D视觉驱动 * TTSA_3D：3D语音驱动 * FLEXUS_2D：FLEXUS版本资源
	BusinessType *ResourceUsageInfoBusinessType `json:"business_type,omitempty"`

	// 子资源类型。
	SubResourceType *string `json:"sub_resource_type,omitempty"`

	// 是否子资源。子资源描述的是子资源的数量和单位
	IsSubResource *bool `json:"is_sub_resource,omitempty"`

	// 计费类型。 * PERIODIC: 包周期 * ONE_TIME：一次性 * ON_DEMAND：按需
	ChargingMode *ResourceUsageInfoChargingMode `json:"charging_mode,omitempty"`

	// 资源来源。 * PURCHASED: 购买 * SP_ALLOCATED：SP分配 * ADMIN_ALLOCATED：系统管理员分配
	ResourceSource *string `json:"resource_source,omitempty"`

	// 总量
	Amount *float32 `json:"amount,omitempty"`

	// 使用量
	Usage *float32 `json:"usage,omitempty"`

	// 单位。 * NUM：个数(形象/声音) * MIN：分钟（视频制作） * HOUR：小时 （直播） * CHANNEL：路（直播/交互） * GB：GB(资产管理) * MILLION_WORDS：百万字 * TEN_THOUSAND_WORDS：万字 * TIME：次
	Unit *ResourceUsageInfoUnit `json:"unit,omitempty"`
}

func (o ResourceUsageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceUsageInfo struct{}"
	}

	return strings.Join([]string{"ResourceUsageInfo", string(data)}, " ")
}

type ResourceUsageInfoBusinessType struct {
	value string
}

type ResourceUsageInfoBusinessTypeEnum struct {
	VOICE_CLONE       ResourceUsageInfoBusinessType
	SYNTHETICS_SOUND  ResourceUsageInfoBusinessType
	ASSET_MANAGER     ResourceUsageInfoBusinessType
	MODELING_2_D      ResourceUsageInfoBusinessType
	LIVE_2_D          ResourceUsageInfoBusinessType
	VIDEO_2_D         ResourceUsageInfoBusinessType
	CHAT_2_D          ResourceUsageInfoBusinessType
	BUSINESS_CARD_2_D ResourceUsageInfoBusinessType
	PICTURE_2_D       ResourceUsageInfoBusinessType
	MODELING_3_D      ResourceUsageInfoBusinessType
	VDS_3_D           ResourceUsageInfoBusinessType
	TTSA_3_D          ResourceUsageInfoBusinessType
	FLEXUS_2_D        ResourceUsageInfoBusinessType
}

func GetResourceUsageInfoBusinessTypeEnum() ResourceUsageInfoBusinessTypeEnum {
	return ResourceUsageInfoBusinessTypeEnum{
		VOICE_CLONE: ResourceUsageInfoBusinessType{
			value: "VOICE_CLONE",
		},
		SYNTHETICS_SOUND: ResourceUsageInfoBusinessType{
			value: "SYNTHETICS_SOUND",
		},
		ASSET_MANAGER: ResourceUsageInfoBusinessType{
			value: "ASSET_MANAGER",
		},
		MODELING_2_D: ResourceUsageInfoBusinessType{
			value: "MODELING_2D",
		},
		LIVE_2_D: ResourceUsageInfoBusinessType{
			value: "LIVE_2D",
		},
		VIDEO_2_D: ResourceUsageInfoBusinessType{
			value: "VIDEO_2D",
		},
		CHAT_2_D: ResourceUsageInfoBusinessType{
			value: "CHAT_2D",
		},
		BUSINESS_CARD_2_D: ResourceUsageInfoBusinessType{
			value: "BUSINESS_CARD_2D",
		},
		PICTURE_2_D: ResourceUsageInfoBusinessType{
			value: "PICTURE_2D",
		},
		MODELING_3_D: ResourceUsageInfoBusinessType{
			value: "MODELING_3D",
		},
		VDS_3_D: ResourceUsageInfoBusinessType{
			value: "VDS_3D",
		},
		TTSA_3_D: ResourceUsageInfoBusinessType{
			value: "TTSA_3D",
		},
		FLEXUS_2_D: ResourceUsageInfoBusinessType{
			value: "FLEXUS_2D",
		},
	}
}

func (c ResourceUsageInfoBusinessType) Value() string {
	return c.value
}

func (c ResourceUsageInfoBusinessType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResourceUsageInfoBusinessType) UnmarshalJSON(b []byte) error {
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

type ResourceUsageInfoChargingMode struct {
	value string
}

type ResourceUsageInfoChargingModeEnum struct {
	PERIODIC  ResourceUsageInfoChargingMode
	ONE_TIME  ResourceUsageInfoChargingMode
	ON_DEMAND ResourceUsageInfoChargingMode
}

func GetResourceUsageInfoChargingModeEnum() ResourceUsageInfoChargingModeEnum {
	return ResourceUsageInfoChargingModeEnum{
		PERIODIC: ResourceUsageInfoChargingMode{
			value: "PERIODIC",
		},
		ONE_TIME: ResourceUsageInfoChargingMode{
			value: "ONE_TIME",
		},
		ON_DEMAND: ResourceUsageInfoChargingMode{
			value: "ON_DEMAND",
		},
	}
}

func (c ResourceUsageInfoChargingMode) Value() string {
	return c.value
}

func (c ResourceUsageInfoChargingMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResourceUsageInfoChargingMode) UnmarshalJSON(b []byte) error {
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

type ResourceUsageInfoUnit struct {
	value string
}

type ResourceUsageInfoUnitEnum struct {
	NUM                ResourceUsageInfoUnit
	MIN                ResourceUsageInfoUnit
	HOUR               ResourceUsageInfoUnit
	CHANNEL            ResourceUsageInfoUnit
	GB                 ResourceUsageInfoUnit
	MILLION_WORDS      ResourceUsageInfoUnit
	TEN_THOUSAND_WORDS ResourceUsageInfoUnit
	TIME               ResourceUsageInfoUnit
}

func GetResourceUsageInfoUnitEnum() ResourceUsageInfoUnitEnum {
	return ResourceUsageInfoUnitEnum{
		NUM: ResourceUsageInfoUnit{
			value: "NUM",
		},
		MIN: ResourceUsageInfoUnit{
			value: "MIN",
		},
		HOUR: ResourceUsageInfoUnit{
			value: "HOUR",
		},
		CHANNEL: ResourceUsageInfoUnit{
			value: "CHANNEL",
		},
		GB: ResourceUsageInfoUnit{
			value: "GB",
		},
		MILLION_WORDS: ResourceUsageInfoUnit{
			value: "MILLION_WORDS",
		},
		TEN_THOUSAND_WORDS: ResourceUsageInfoUnit{
			value: "TEN_THOUSAND_WORDS",
		},
		TIME: ResourceUsageInfoUnit{
			value: "TIME",
		},
	}
}

func (c ResourceUsageInfoUnit) Value() string {
	return c.value
}

func (c ResourceUsageInfoUnit) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResourceUsageInfoUnit) UnmarshalJSON(b []byte) error {
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
