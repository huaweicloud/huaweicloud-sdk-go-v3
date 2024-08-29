package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BackgroundConfigInfo 背景配置。
type BackgroundConfigInfo struct {

	// 背景类型。 - IMAGE：图片背景，指定图片用作分身数字人背景。 - COLOR：纯色背景，指定颜色RGB值作为分身数字人背景。
	BackgroundType BackgroundConfigInfoBackgroundType `json:"background_type"`

	// 背景文件的URL。 > * 仅直播支持外部URL，其他业务通过资产库查询获取，不支持外部URL。 > * background_type=IMAGE时需要填写。
	BackgroundConfig *string `json:"background_config,omitempty"`

	// 纯色背景的RGB颜色值。 > * background_type=COLOR时需要填写。
	BackgroundColorConfig *string `json:"background_color_config,omitempty"`

	// 背景资产ID。 > * 背景是背景图片时，填图片资产ID，可以从资产库中查询。
	BackgroundAssetId *string `json:"background_asset_id,omitempty"`
}

func (o BackgroundConfigInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackgroundConfigInfo struct{}"
	}

	return strings.Join([]string{"BackgroundConfigInfo", string(data)}, " ")
}

type BackgroundConfigInfoBackgroundType struct {
	value string
}

type BackgroundConfigInfoBackgroundTypeEnum struct {
	IMAGE     BackgroundConfigInfoBackgroundType
	COLOR     BackgroundConfigInfoBackgroundType
	IMAGE_2_D BackgroundConfigInfoBackgroundType
	VIDEO     BackgroundConfigInfoBackgroundType
	AUDIO     BackgroundConfigInfoBackgroundType
}

func GetBackgroundConfigInfoBackgroundTypeEnum() BackgroundConfigInfoBackgroundTypeEnum {
	return BackgroundConfigInfoBackgroundTypeEnum{
		IMAGE: BackgroundConfigInfoBackgroundType{
			value: "IMAGE",
		},
		COLOR: BackgroundConfigInfoBackgroundType{
			value: "COLOR",
		},
		IMAGE_2_D: BackgroundConfigInfoBackgroundType{
			value: "IMAGE_2D",
		},
		VIDEO: BackgroundConfigInfoBackgroundType{
			value: "VIDEO",
		},
		AUDIO: BackgroundConfigInfoBackgroundType{
			value: "AUDIO",
		},
	}
}

func (c BackgroundConfigInfoBackgroundType) Value() string {
	return c.value
}

func (c BackgroundConfigInfoBackgroundType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BackgroundConfigInfoBackgroundType) UnmarshalJSON(b []byte) error {
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
