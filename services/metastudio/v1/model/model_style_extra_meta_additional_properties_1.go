package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StyleExtraMetaAdditionalProperties1 struct {

	// 算法类型枚举，\"CREATE\"还是\"CLASSIFY\"
	Type string `json:"type"`

	// 算法输出所匹配的组件名
	MatchComponent *string `json:"match_component,omitempty"`

	// 算法输出的文件名列表
	Files *[]string `json:"files,omitempty"`

	// 分类算法的标签列表
	ClassifiedTags map[string][]string `json:"classified_tags,omitempty"`
}

func (o StyleExtraMetaAdditionalProperties1) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StyleExtraMetaAdditionalProperties1 struct{}"
	}

	return strings.Join([]string{"StyleExtraMetaAdditionalProperties1", string(data)}, " ")
}
