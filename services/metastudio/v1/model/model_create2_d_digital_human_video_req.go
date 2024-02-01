package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Create2DDigitalHumanVideoReq struct {

	// 剧本ID。 > * 如果填写了script_id，model_asset_id、voice_config、scene_asset_id、video_config、shoot_scripts可以不填，以脚本中的配置为准。 > * 如果填写了script_id，并且同时也填写了model_asset_id、voice_config、scene_asset_id、video_config、shoot_scripts则以本接口中的配置为准。
	ScriptId *string `json:"script_id,omitempty"`

	// 视频生成类型。该参数取值是MODEL时，model_asset_id必填；取值是PICTURE时，human_image必填。 * MODEL：通过分身数字人模型生成视频 * PICTURE： 通过单张照片生成视频 > * 该参数已废弃，照片数字人视频制作使用“创建照片分身数字人视频制作任务”接口。
	VideoMakingType *Create2DDigitalHumanVideoReqVideoMakingType `json:"video_making_type,omitempty"`

	// 分身数字人模型资产ID。
	ModelAssetId *string `json:"model_asset_id,omitempty"`

	// 人物照片，需要Base64编码。照片分辨率不超过1080P。 > * 该参数已废弃，照片数字人视频制作使用“创建照片分身数字人视频制作任务”接口。
	HumanImage *string `json:"human_image,omitempty"`

	VoiceConfig *VoiceConfig `json:"voice_config,omitempty"`

	VideoConfig *VideoConfig `json:"video_config,omitempty"`

	// 拍摄脚本列表。
	ShootScripts *[]ShootScriptItem `json:"shoot_scripts,omitempty"`

	OutputAssetConfig *OutputAssetConfig `json:"output_asset_config,omitempty"`

	BackgroundMusicConfig *BackgroundMusicConfig `json:"background_music_config,omitempty"`

	ReviewConfig *ReviewConfig `json:"review_config,omitempty"`

	CallbackConfig *CallBackConfig `json:"callback_config,omitempty"`
}

func (o Create2DDigitalHumanVideoReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Create2DDigitalHumanVideoReq struct{}"
	}

	return strings.Join([]string{"Create2DDigitalHumanVideoReq", string(data)}, " ")
}

type Create2DDigitalHumanVideoReqVideoMakingType struct {
	value string
}

type Create2DDigitalHumanVideoReqVideoMakingTypeEnum struct {
	MODEL   Create2DDigitalHumanVideoReqVideoMakingType
	PICTURE Create2DDigitalHumanVideoReqVideoMakingType
}

func GetCreate2DDigitalHumanVideoReqVideoMakingTypeEnum() Create2DDigitalHumanVideoReqVideoMakingTypeEnum {
	return Create2DDigitalHumanVideoReqVideoMakingTypeEnum{
		MODEL: Create2DDigitalHumanVideoReqVideoMakingType{
			value: "MODEL",
		},
		PICTURE: Create2DDigitalHumanVideoReqVideoMakingType{
			value: "PICTURE",
		},
	}
}

func (c Create2DDigitalHumanVideoReqVideoMakingType) Value() string {
	return c.value
}

func (c Create2DDigitalHumanVideoReqVideoMakingType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *Create2DDigitalHumanVideoReqVideoMakingType) UnmarshalJSON(b []byte) error {
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
