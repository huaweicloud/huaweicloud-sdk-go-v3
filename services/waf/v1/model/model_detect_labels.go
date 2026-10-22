package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetectLabels 检测标签配置
type DetectLabels struct {

	// 合规检测标签列表
	ComplianceLabels *[]string `json:"compliance_labels,omitempty"`

	// 注入检测标签列表
	InjectLabels *[]string `json:"inject_labels,omitempty"`

	// 豁免词库标签列表
	ImmunityLexiconLabels *[]string `json:"immunity_lexicon_labels,omitempty"`

	// 自定义词库检测标签列表
	LexiconLabels *[]string `json:"lexicon_labels,omitempty"`

	// 内置词库检测标签列表
	LexiconInternalLabels *[]string `json:"lexicon_internal_labels,omitempty"`

	// 支持检测语言标签列表
	SupportedLanguageLabels *[]string `json:"supported_language_labels,omitempty"`

	// 敏感词标签列表
	PiiLabels *[]string `json:"pii_labels,omitempty"`
}

func (o DetectLabels) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetectLabels struct{}"
	}

	return strings.Join([]string{"DetectLabels", string(data)}, " ")
}
