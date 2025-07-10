package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type SystemProperty struct {

	// **参数解释**： 操作。 **约束限制**： 系统属性仅为系统设置，普通用户无法修改。 **取值范围**： * ADD：增加 * DELETE：删除  **默认取值**： 不涉及
	Action *SystemPropertyAction `json:"action,omitempty"`

	// **参数解释**： 系统属性条目。 **约束限制**： 系统属性仅为系统设置，普通用户无法修改。 **取值范围**： 公共资产属性： * BACKGROUND_IMG：视频制作的背景图片。value设置成Yes * CREATED_BY_PLATFORM: 是否平台生成。 * BACKGROUND_SCENE：视频制作的2D背景场景。value可选Horizontal（横屏）或者Vertical（竖屏）。 * MEITUAN_MATERIAL_APPROVED: 美团平台已审核标识，value设置成YES。  分身数字人资产属性： * MATERIAL_IMG：素材图片，用作图片图层。value设置成Yes，否则控制台视频制作、直播等界面的贴图区域，将无法看到此图片。 * MATERIAL_VIDEO：素材视频，用作视频图层。value设置成Yes，否则控制台视频制作、直播等界面的视频区域，将无法看到此视频。 * DIGITAL_HUMAN_2D_VIDEO：分身数字人视频。 * BUSINESS_CARD_VIDEO：名片视频。 * BUSSINESS_CARD_VIDEO：名片视频(过期) * PHOTO_VIDEO：照片数字人视频。  **默认取值**： 不涉及
	Key *SystemPropertyKey `json:"key,omitempty"`

	// **参数解释**： 系统属性属性值。 **约束限制**： 系统属性仅为系统设置，普通用户无法修改。 **取值范围**： 字符长度1-1024位 **默认取值** 不涉及
	Value *string `json:"value,omitempty"`
}

func (o SystemProperty) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SystemProperty struct{}"
	}

	return strings.Join([]string{"SystemProperty", string(data)}, " ")
}

type SystemPropertyAction struct {
	value string
}

type SystemPropertyActionEnum struct {
	ADD    SystemPropertyAction
	DELETE SystemPropertyAction
}

func GetSystemPropertyActionEnum() SystemPropertyActionEnum {
	return SystemPropertyActionEnum{
		ADD: SystemPropertyAction{
			value: "ADD",
		},
		DELETE: SystemPropertyAction{
			value: "DELETE",
		},
	}
}

func (c SystemPropertyAction) Value() string {
	return c.value
}

func (c SystemPropertyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SystemPropertyAction) UnmarshalJSON(b []byte) error {
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

type SystemPropertyKey struct {
	value string
}

type SystemPropertyKeyEnum struct {
	STYLE_ID                  SystemPropertyKey
	DH_ID                     SystemPropertyKey
	PLATFORM_AVAILABLE        SystemPropertyKey
	RENDER_ENGINE             SystemPropertyKey
	BACKGROUND_IMG            SystemPropertyKey
	BACKGROUND_SCENE          SystemPropertyKey
	CREATED_BY_PLATFORM       SystemPropertyKey
	MATERIAL_IMG              SystemPropertyKey
	MATERIAL_VIDEO            SystemPropertyKey
	DIGITAL_HUMAN_2_D_VIDEO   SystemPropertyKey
	DIGITAL_HUMAN_3_D_VIDEO   SystemPropertyKey
	BUSSINESS_CARD_VIDEO      SystemPropertyKey
	BUSINESS_CARD_VIDEO       SystemPropertyKey
	PHOTO_VIDEO               SystemPropertyKey
	TO_BE_TRANSLATED_VIDEO    SystemPropertyKey
	TRANSLATED_VIDEO          SystemPropertyKey
	LAYER_CONFIG_ENABLE       SystemPropertyKey
	MEITUAN_MATERIAL_APPROVED SystemPropertyKey
}

func GetSystemPropertyKeyEnum() SystemPropertyKeyEnum {
	return SystemPropertyKeyEnum{
		STYLE_ID: SystemPropertyKey{
			value: "STYLE_ID",
		},
		DH_ID: SystemPropertyKey{
			value: "DH_ID",
		},
		PLATFORM_AVAILABLE: SystemPropertyKey{
			value: "PLATFORM_AVAILABLE",
		},
		RENDER_ENGINE: SystemPropertyKey{
			value: "RENDER_ENGINE",
		},
		BACKGROUND_IMG: SystemPropertyKey{
			value: "BACKGROUND_IMG",
		},
		BACKGROUND_SCENE: SystemPropertyKey{
			value: "BACKGROUND_SCENE",
		},
		CREATED_BY_PLATFORM: SystemPropertyKey{
			value: "CREATED_BY_PLATFORM",
		},
		MATERIAL_IMG: SystemPropertyKey{
			value: "MATERIAL_IMG",
		},
		MATERIAL_VIDEO: SystemPropertyKey{
			value: "MATERIAL_VIDEO",
		},
		DIGITAL_HUMAN_2_D_VIDEO: SystemPropertyKey{
			value: "DIGITAL_HUMAN_2D_VIDEO",
		},
		DIGITAL_HUMAN_3_D_VIDEO: SystemPropertyKey{
			value: "DIGITAL_HUMAN_3D_VIDEO",
		},
		BUSSINESS_CARD_VIDEO: SystemPropertyKey{
			value: "BUSSINESS_CARD_VIDEO",
		},
		BUSINESS_CARD_VIDEO: SystemPropertyKey{
			value: "BUSINESS_CARD_VIDEO",
		},
		PHOTO_VIDEO: SystemPropertyKey{
			value: "PHOTO_VIDEO",
		},
		TO_BE_TRANSLATED_VIDEO: SystemPropertyKey{
			value: "TO_BE_TRANSLATED_VIDEO",
		},
		TRANSLATED_VIDEO: SystemPropertyKey{
			value: "TRANSLATED_VIDEO",
		},
		LAYER_CONFIG_ENABLE: SystemPropertyKey{
			value: "LAYER_CONFIG_ENABLE",
		},
		MEITUAN_MATERIAL_APPROVED: SystemPropertyKey{
			value: "MEITUAN_MATERIAL_APPROVED",
		},
	}
}

func (c SystemPropertyKey) Value() string {
	return c.value
}

func (c SystemPropertyKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SystemPropertyKey) UnmarshalJSON(b []byte) error {
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
