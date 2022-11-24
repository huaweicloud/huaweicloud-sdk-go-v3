package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TextDetectionResultDetail struct {

	// 审核结果是否通过。  block：包含敏感信息，不通过  pass：不包含敏感信息，通过  review：需要人工复检
	Suggestion *string `json:"suggestion,omitempty"`

	// 检测结果的标签。  支持label列表如下：  politics: 涉政  terrorism: 暴恐  porn: 色情  ban: 违禁  abuse: 辱骂  ad: 广告  ad_law: 广告法  meaningless: ⽆意义  customized：自定义（命中自定义词库中的关键词）
	Label *string `json:"label,omitempty"`

	// 置信度，取值范围 0-1，值越⼤，可信度越⾼。
	Confidence *float32 `json:"confidence,omitempty"`

	// 命中的风险片段信息，如果命中了语义算法模型，则会返回一个空的列表。
	Segments *[]SegmentResult `json:"segments,omitempty"`
}

func (o TextDetectionResultDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TextDetectionResultDetail struct{}"
	}

	return strings.Join([]string{"TextDetectionResultDetail", string(data)}, " ")
}
