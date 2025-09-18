package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PipelineTemplateSimpleVoStages struct {

	// **参数解释**： 阶段名称。 **取值范围**： 仅支持输入中文、大小写英文字母、数字、'-'、'_'、','、';'、':'、'.'、'/'、'('、')'、'（'、'）'及空格，其中空格不可在名称开头或结尾使用，且长度为[1,128]个字符。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 阶段顺序。 **取值范围**： 大于等于0。
	Sequence *int32 `json:"sequence,omitempty"`
}

func (o PipelineTemplateSimpleVoStages) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipelineTemplateSimpleVoStages struct{}"
	}

	return strings.Join([]string{"PipelineTemplateSimpleVoStages", string(data)}, " ")
}
