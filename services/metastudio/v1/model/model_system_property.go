package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type SystemProperty struct {

	// 操作。 - ADD：增加 - DELETE：删除
	Action *SystemPropertyAction `json:"action,omitempty"`

	// 系统属性。 * STYLE_ID：风格Id。 * DH_ID：数字人ID(尚未启用)。 * PLATFORM_AVAILABLE：是否平台可用(尚未启用)。 * RENDER_ENGINE：引擎类型。value可选UE或MetaEngine。 * BACKGROUND_IMG：视频制作的2D背景图片。value设置成Yes * BACKGROUND_SCENE：视频制作的2D背景场景。value可选Horizontal（横屏）或者Vertical（竖屏） * CREATED_BY_PLATFORM： 是否平台生成
	Key *SystemPropertyKey `json:"key,omitempty"`

	// 属性值。
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
	STYLE_ID            SystemPropertyKey
	DH_ID               SystemPropertyKey
	PLATFORM_AVAILABLE  SystemPropertyKey
	RENDER_ENGINE       SystemPropertyKey
	BACKGROUND_IMG      SystemPropertyKey
	BACKGROUND_SCENE    SystemPropertyKey
	CREATED_BY_PLATFORM SystemPropertyKey
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
