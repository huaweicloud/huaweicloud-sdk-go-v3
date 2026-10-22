package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetectLabelsInfo 检测标签详情
type DetectLabelsInfo struct {

	// 合规检测标签详情列表
	ComplianceLabelsInfo *[]LabelInfo `json:"compliance_labels_info,omitempty"`

	// 注入检测标签详情列表
	InjectLabelsInfo *[]LabelInfo `json:"inject_labels_info,omitempty"`

	// 豁免词库标签详情列表
	ImmunityLexiconLabelsInfo *[]LabelInfo `json:"immunity_lexicon_labels_info,omitempty"`

	// 自定义词库检测标签详情列表
	LexiconLabelsInfo *[]LabelInfo `json:"lexicon_labels_info,omitempty"`

	// 内置词库检测标签详情列表
	LexiconInternalLabelsInfo *[]LabelInfo `json:"lexicon_internal_labels_info,omitempty"`

	// 语言检测标签信息列表
	SupportedLanguageLabels *[]LabelInfo `json:"supported_language_labels,omitempty"`

	// 敏感词的标签详情列表
	PiiLabelsInfo *[]LabelInfo `json:"pii_labels_info,omitempty"`
}

func (o DetectLabelsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetectLabelsInfo struct{}"
	}

	return strings.Join([]string{"DetectLabelsInfo", string(data)}, " ")
}
