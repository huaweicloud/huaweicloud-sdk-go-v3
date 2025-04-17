package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AdditionalManifest struct {

	// 索引后缀名，后缀名仅支持数字，字母、下划线、中划线。
	ManifestNameModifier string `json:"manifest_name_modifier"`

	// 选择流名列表，最多支持9路流。
	SelectedOutputs []string `json:"selected_outputs"`
}

func (o AdditionalManifest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AdditionalManifest struct{}"
	}

	return strings.Join([]string{"AdditionalManifest", string(data)}, " ")
}
