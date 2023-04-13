package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 场景组件信息。
type SceneComponentInfo struct {

	// 组件索引。
	Index *int32 `json:"index,omitempty"`

	// 组件名称。
	ComponentName string `json:"component_name"`

	// 组件类型。 * CAMERA： 摄像机 * PANEL： 屏幕 * LIGHT： 灯光
	ComponentType SceneComponentInfoComponentType `json:"component_type"`

	// 组件描述。
	ComponentDesc *string `json:"component_desc,omitempty"`
}

func (o SceneComponentInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SceneComponentInfo struct{}"
	}

	return strings.Join([]string{"SceneComponentInfo", string(data)}, " ")
}

type SceneComponentInfoComponentType struct {
	value string
}

type SceneComponentInfoComponentTypeEnum struct {
	CAMERA SceneComponentInfoComponentType
	PANEL  SceneComponentInfoComponentType
	LIGHT  SceneComponentInfoComponentType
}

func GetSceneComponentInfoComponentTypeEnum() SceneComponentInfoComponentTypeEnum {
	return SceneComponentInfoComponentTypeEnum{
		CAMERA: SceneComponentInfoComponentType{
			value: "CAMERA",
		},
		PANEL: SceneComponentInfoComponentType{
			value: "PANEL",
		},
		LIGHT: SceneComponentInfoComponentType{
			value: "LIGHT",
		},
	}
}

func (c SceneComponentInfoComponentType) Value() string {
	return c.value
}

func (c SceneComponentInfoComponentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SceneComponentInfoComponentType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
