package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StyleExtraMetaAdditionalProperties 可替换组件
type StyleExtraMetaAdditionalProperties struct {

	// 图标url
	Icon *string `json:"icon,omitempty"`

	// 展示标签 {\"cn\": \"xxxxx\",\"en\":\"xxxxx\"}
	Label map[string]string `json:"label,omitempty"`

	// 是否使用算法输出文件，针对生成算法
	UseAlgFile *bool `json:"use_alg_file,omitempty"`

	// 所属算法标签属性值，针对分类算法
	AlgorithmClassifyTag map[string]string `json:"algorithm_classify_tag,omitempty"`

	// 当前使用的身体骨骼
	ModelBodyStyle *string `json:"model_body_style,omitempty"`

	McAsset *EngineAssetInfo `json:"mc_asset,omitempty"`

	UeAsset *EngineAssetInfo `json:"ue_asset,omitempty"`
}

func (o StyleExtraMetaAdditionalProperties) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StyleExtraMetaAdditionalProperties struct{}"
	}

	return strings.Join([]string{"StyleExtraMetaAdditionalProperties", string(data)}, " ")
}
