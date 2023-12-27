package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BackgroundConfigInfo 背景配置。
type BackgroundConfigInfo struct {

	// 背景类型。 - IMAGE：图片，用于3D数字人演示素材讲解模式的图片或分身数字背景图片 - IMAGE_2D：图片，用于3D数字人主播播报模式的2D场景背景图片 - VIDEO：视频 - AUDIO：音频 > * 分身数字人视频制作仅支持IMAGE
	BackgroundType BackgroundConfigInfoBackgroundType `json:"background_type"`

	// 背景标题。 > * 分身数字人视频制作此参数不生效。
	BackgroundTitle *string `json:"background_title,omitempty"`

	HumanPosition2d *HumanPosition2D `json:"human_position_2d,omitempty"`

	HumanSize2d *HumanSize2D `json:"human_size_2d,omitempty"`

	// 视频文件封面图片的下载URL。  演示素材为视频时有效。 > * 分身数字人视频制作此参数不生效。
	BackgroundCoverUrl *string `json:"background_cover_url,omitempty"`

	// 背景文件的URL。 > * 通过资产库查询获取，不支持外部URL。
	BackgroundConfig string `json:"background_config"`

	// 背景资产ID。 > * 背景是背景图片时，填图片资产ID。
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
	IMAGE_2_D BackgroundConfigInfoBackgroundType
	VIDEO     BackgroundConfigInfoBackgroundType
	AUDIO     BackgroundConfigInfoBackgroundType
}

func GetBackgroundConfigInfoBackgroundTypeEnum() BackgroundConfigInfoBackgroundTypeEnum {
	return BackgroundConfigInfoBackgroundTypeEnum{
		IMAGE: BackgroundConfigInfoBackgroundType{
			value: "IMAGE",
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
