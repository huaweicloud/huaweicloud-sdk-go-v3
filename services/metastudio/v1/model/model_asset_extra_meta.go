package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 资产元数据。根据资产类型选择其中一个填写。
type AssetExtraMeta struct {
	HumanModelMeta *HumanModelAssetMeta `json:"human_model_meta,omitempty"`

	VoiceModelMeta *VoiceModelAssetMeta `json:"voice_model_meta,omitempty"`

	PptMeta *PptAssetMeta `json:"ppt_meta,omitempty"`

	AnimationMeta *AnimationAssetMeta `json:"animation_meta,omitempty"`

	SceneMeta *SceneAssetMeta `json:"scene_meta,omitempty"`

	MaterialMeta *MaterialAssetMeta `json:"material_meta,omitempty"`
}

func (o AssetExtraMeta) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssetExtraMeta struct{}"
	}

	return strings.Join([]string{"AssetExtraMeta", string(data)}, " ")
}
