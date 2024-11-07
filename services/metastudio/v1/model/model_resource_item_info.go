package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ResourceItemInfo 资源使用信息
type ResourceItemInfo struct {

	// 资源ID。
	ResourceId *string `json:"resource_id,omitempty"`

	// cloudServiceId的订单id。 云服务购买清单等场景必填（purchaseMode取值为3、4），每个CloudService生成一个订单；此场景如果为空，则报错。 其他场景，为空。
	OrderId *string `json:"order_id,omitempty"`

	// 资源截止时间。格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"
	ResourceExpireTime *string `json:"resource_expire_time,omitempty"`

	// 资源类型。详见[资源类型](metastudio_02_0042.xml)。
	ResourceType *string `json:"resource_type,omitempty"`

	// 业务类型。 * VOICE_CLONE：声音制作 * SYNTHETICS_SOUND：声音合成 * ASSET_MANAGER：资产管理 * MODELING_2D：形象制作 * LIVE_2D：分身数字人视频直播 * VIDEO_2D：分身数字人视频制作 * CHAT_2D：分身数字人智能交互 * BUSINESS_CARD_2D：分身数字人名片 * PICTURE_2D：照片数字人视频 * MODELING_3D：3D照片建模 * VDS_3D：3D视觉驱动 * TTSA_3D：3D语音驱动 * FLEXUS_2D：flexus版本资源
	BusinessType *ResourceItemInfoBusinessType `json:"business_type,omitempty"`

	// 子资源类型。当前只有flexus套餐包存在该字段 * voice_clone_flexus: 语音克隆Flexus版 * modeling_count_2d_model_flexus: 分身数字人形象制作Flexus版 * video_time_flexus_2d_model: 分身数字人Flexus版本视频制作
	SubResourceType *string `json:"sub_resource_type,omitempty"`

	// 是否子资源。子资源描述的是子资源的数量和单位
	IsSubResource *bool `json:"is_sub_resource,omitempty"`

	// 计费类型。 * PERIODIC: 包周期 * ONE_TIME：一次性 * ON_DEMAND：按需
	ChargingMode *ResourceItemInfoChargingMode `json:"charging_mode,omitempty"`

	// 资源来源。 * PURCHASED: 购买 * SP_ALLOCATED：SP分配 * ADMIN_ALLOCATED：系统管理员分配
	ResourceSource *string `json:"resource_source,omitempty"`

	// 总量
	Amount float32 `json:"amount,omitempty"`

	// 使用量
	Usage float32 `json:"usage,omitempty"`

	// 单位。 * NUM：个数(形象/声音) * MIN：分钟（视频制作） * HOUR：小时 （直播） * CHANNEL：路（直播/交互） * GB：GB(资产管理) * MILLION_WORDS：百万字 * TEN_THOUSAND_WORDS：万字 * TIME：次
	Unit *ResourceItemInfoUnit `json:"unit,omitempty"`
}

func (o ResourceItemInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceItemInfo struct{}"
	}

	return strings.Join([]string{"ResourceItemInfo", string(data)}, " ")
}

type ResourceItemInfoBusinessType struct {
	value string
}

type ResourceItemInfoBusinessTypeEnum struct {
	VOICE_CLONE       ResourceItemInfoBusinessType
	SYNTHETICS_SOUND  ResourceItemInfoBusinessType
	ASSET_MANAGER     ResourceItemInfoBusinessType
	MODELING_2_D      ResourceItemInfoBusinessType
	LIVE_2_D          ResourceItemInfoBusinessType
	VIDEO_2_D         ResourceItemInfoBusinessType
	CHAT_2_D          ResourceItemInfoBusinessType
	BUSINESS_CARD_2_D ResourceItemInfoBusinessType
	PICTURE_2_D       ResourceItemInfoBusinessType
	MODELING_3_D      ResourceItemInfoBusinessType
	VDS_3_D           ResourceItemInfoBusinessType
	TTSA_3_D          ResourceItemInfoBusinessType
	FLEXUS_2_D        ResourceItemInfoBusinessType
}

func GetResourceItemInfoBusinessTypeEnum() ResourceItemInfoBusinessTypeEnum {
	return ResourceItemInfoBusinessTypeEnum{
		VOICE_CLONE: ResourceItemInfoBusinessType{
			value: "VOICE_CLONE",
		},
		SYNTHETICS_SOUND: ResourceItemInfoBusinessType{
			value: "SYNTHETICS_SOUND",
		},
		ASSET_MANAGER: ResourceItemInfoBusinessType{
			value: "ASSET_MANAGER",
		},
		MODELING_2_D: ResourceItemInfoBusinessType{
			value: "MODELING_2D",
		},
		LIVE_2_D: ResourceItemInfoBusinessType{
			value: "LIVE_2D",
		},
		VIDEO_2_D: ResourceItemInfoBusinessType{
			value: "VIDEO_2D",
		},
		CHAT_2_D: ResourceItemInfoBusinessType{
			value: "CHAT_2D",
		},
		BUSINESS_CARD_2_D: ResourceItemInfoBusinessType{
			value: "BUSINESS_CARD_2D",
		},
		PICTURE_2_D: ResourceItemInfoBusinessType{
			value: "PICTURE_2D",
		},
		MODELING_3_D: ResourceItemInfoBusinessType{
			value: "MODELING_3D",
		},
		VDS_3_D: ResourceItemInfoBusinessType{
			value: "VDS_3D",
		},
		TTSA_3_D: ResourceItemInfoBusinessType{
			value: "TTSA_3D",
		},
		FLEXUS_2_D: ResourceItemInfoBusinessType{
			value: "FLEXUS_2D",
		},
	}
}

func (c ResourceItemInfoBusinessType) Value() string {
	return c.value
}

func (c ResourceItemInfoBusinessType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResourceItemInfoBusinessType) UnmarshalJSON(b []byte) error {
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

type ResourceItemInfoChargingMode struct {
	value string
}

type ResourceItemInfoChargingModeEnum struct {
	PERIODIC  ResourceItemInfoChargingMode
	ONE_TIME  ResourceItemInfoChargingMode
	ON_DEMAND ResourceItemInfoChargingMode
}

func GetResourceItemInfoChargingModeEnum() ResourceItemInfoChargingModeEnum {
	return ResourceItemInfoChargingModeEnum{
		PERIODIC: ResourceItemInfoChargingMode{
			value: "PERIODIC",
		},
		ONE_TIME: ResourceItemInfoChargingMode{
			value: "ONE_TIME",
		},
		ON_DEMAND: ResourceItemInfoChargingMode{
			value: "ON_DEMAND",
		},
	}
}

func (c ResourceItemInfoChargingMode) Value() string {
	return c.value
}

func (c ResourceItemInfoChargingMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResourceItemInfoChargingMode) UnmarshalJSON(b []byte) error {
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

type ResourceItemInfoUnit struct {
	value string
}

type ResourceItemInfoUnitEnum struct {
	NUM                ResourceItemInfoUnit
	MIN                ResourceItemInfoUnit
	HOUR               ResourceItemInfoUnit
	CHANNEL            ResourceItemInfoUnit
	GB                 ResourceItemInfoUnit
	MILLION_WORDS      ResourceItemInfoUnit
	TEN_THOUSAND_WORDS ResourceItemInfoUnit
	TIME               ResourceItemInfoUnit
}

func GetResourceItemInfoUnitEnum() ResourceItemInfoUnitEnum {
	return ResourceItemInfoUnitEnum{
		NUM: ResourceItemInfoUnit{
			value: "NUM",
		},
		MIN: ResourceItemInfoUnit{
			value: "MIN",
		},
		HOUR: ResourceItemInfoUnit{
			value: "HOUR",
		},
		CHANNEL: ResourceItemInfoUnit{
			value: "CHANNEL",
		},
		GB: ResourceItemInfoUnit{
			value: "GB",
		},
		MILLION_WORDS: ResourceItemInfoUnit{
			value: "MILLION_WORDS",
		},
		TEN_THOUSAND_WORDS: ResourceItemInfoUnit{
			value: "TEN_THOUSAND_WORDS",
		},
		TIME: ResourceItemInfoUnit{
			value: "TIME",
		},
	}
}

func (c ResourceItemInfoUnit) Value() string {
	return c.value
}

func (c ResourceItemInfoUnit) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResourceItemInfoUnit) UnmarshalJSON(b []byte) error {
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
