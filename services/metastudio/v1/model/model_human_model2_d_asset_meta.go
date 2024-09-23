package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// HumanModel2DAssetMeta 分身模型元数据
type HumanModel2DAssetMeta struct {

	// **参数解释**： 分身数字人的动作是否可编辑 **约束限制**： 不涉及 **取值范围**： * true: 动作可编辑 * false: 动作不可编辑
	IsActionEditable *bool `json:"is_action_editable,omitempty"`

	// **参数解释**： 是否是实景分身数字人 **约束限制**： 实景分身数字人不做背景替换。 **取值范围**： * true: 实景分身数字人 * false: 普通分身数字人，不带背景。
	IsRealBackground *bool `json:"is_real_background,omitempty"`

	// **参数解释**： 是否支持直播 **约束限制**： 不涉及 **取值范围**： * true: 支持直播 * false: 不支持直播。
	SupportLive *bool `json:"support_live,omitempty"`

	// **参数解释**： 分身数字人模型版本 **约束限制**： 不涉及 **取值范围**： * V2: V2版本模型 * V3：V3版本模型 * V3_2：V3.2版本模型
	ModelVersion *HumanModel2DAssetMetaModelVersion `json:"model_version,omitempty"`

	// **参数解释**： 分身数字人模型分辨率。 **约束限制**： 不涉及 **取值范围**： * 1080P：1080P。支持1080P及720P的视频输出。 * 4K：4K。支持4K、1080P及720P的视频输出。
	ModelResolution *string `json:"model_resolution,omitempty"`

	// **参数解释**： 已执行编译任务设备类型列表。 **约束限制**： 支持走动的数字人，当前仅用于视频制作，不能用于直播和智能交互 **取值范围**： 设备名称列表最多16个。 设备名称字符长度1-64位。 **默认取值**： false
	DeviceNames *[]string `json:"device_names,omitempty"`

	// 分身数字人是否带原子动作库。 > * 带原子动作库的分身数字人可做动作编排。
	IsWithActionLibrary *bool `json:"is_with_action_library,omitempty"`

	// 动作标签映射。
	ActionTagMap *[]ActionTagInfo `json:"action_tag_map,omitempty"`

	// 是否是Flexus版本分身数字人。
	IsFlexus *bool `json:"is_flexus,omitempty"`
}

func (o HumanModel2DAssetMeta) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HumanModel2DAssetMeta struct{}"
	}

	return strings.Join([]string{"HumanModel2DAssetMeta", string(data)}, " ")
}

type HumanModel2DAssetMetaModelVersion struct {
	value string
}

type HumanModel2DAssetMetaModelVersionEnum struct {
	V2   HumanModel2DAssetMetaModelVersion
	V3   HumanModel2DAssetMetaModelVersion
	V3_2 HumanModel2DAssetMetaModelVersion
}

func GetHumanModel2DAssetMetaModelVersionEnum() HumanModel2DAssetMetaModelVersionEnum {
	return HumanModel2DAssetMetaModelVersionEnum{
		V2: HumanModel2DAssetMetaModelVersion{
			value: "V2",
		},
		V3: HumanModel2DAssetMetaModelVersion{
			value: "V3",
		},
		V3_2: HumanModel2DAssetMetaModelVersion{
			value: "V3_2",
		},
	}
}

func (c HumanModel2DAssetMetaModelVersion) Value() string {
	return c.value
}

func (c HumanModel2DAssetMetaModelVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *HumanModel2DAssetMetaModelVersion) UnmarshalJSON(b []byte) error {
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
