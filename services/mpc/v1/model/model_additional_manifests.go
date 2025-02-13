package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AdditionalManifests struct {

	// 定制的索引后缀名
	ManifestNameModifier *string `json:"manifest_name_modifier,omitempty"`

	// 选择的流名称
	SelectedOutputs *[]string `json:"selected_outputs,omitempty"`
}

func (o AdditionalManifests) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AdditionalManifests struct{}"
	}

	return strings.Join([]string{"AdditionalManifests", string(data)}, " ")
}
