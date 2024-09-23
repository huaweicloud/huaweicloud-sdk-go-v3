package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssetExtraMeta **参数解释**： 资产额外元数据。  **约束限制**： 根据asset_type选择对应结构填写，填写其他不匹配的结构会被忽略。 * HUMAN_MODEL: 填写human_model_meta * HUMAN_MODEL_2D: 填写human_model_2d_meta * SCENE：填写scene_meta * ANIMATION：填写animation_meta * MATERIAL：填写material_meta * VOICE_MODEL：填写voice_model_meta * VIDEO：填写video_meta * IMAGE：填写image_meta * PPT：填写ppt_meta * AUDIO: 填写audio_meta
type AssetExtraMeta struct {
	HumanModelMeta *HumanModelAssetMeta `json:"human_model_meta,omitempty"`

	VoiceModelMeta *VoiceModelAssetMeta `json:"voice_model_meta,omitempty"`

	PptMeta *PptAssetMeta `json:"ppt_meta,omitempty"`

	AnimationMeta *AnimationAssetMeta `json:"animation_meta,omitempty"`

	SceneMeta *SceneAssetMeta `json:"scene_meta,omitempty"`

	MaterialMeta *MaterialAssetMeta `json:"material_meta,omitempty"`

	HumanModel2dMeta *HumanModel2DAssetMeta `json:"human_model_2d_meta,omitempty"`

	ImageMeta *ImageAssetMeta `json:"image_meta,omitempty"`

	VideoMeta *VideoAssetMeta `json:"video_meta,omitempty"`

	AudioMeta *AudioAssetMeta `json:"audio_meta,omitempty"`
}

func (o AssetExtraMeta) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssetExtraMeta struct{}"
	}

	return strings.Join([]string{"AssetExtraMeta", string(data)}, " ")
}
