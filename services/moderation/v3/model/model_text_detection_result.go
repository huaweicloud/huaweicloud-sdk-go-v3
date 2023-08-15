package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TextDetectionResult struct {

	// 审核结果是否通过。  block：包含敏感信息，不通过  pass：不包含敏感信息，通过  review：需要人工复检
	Suggestion *string `json:"suggestion,omitempty"`

	// 检测结果的标签。  支持label列表如下：  politics: 涉政  terrorism: 暴恐  porn: 色情  ban: 违禁  abuse: 辱骂  ad: 广告  ad_law: 广告法  meaningless: ⽆意义  customized：自定义（命中自定义词库中的关键词）
	Label *string `json:"label,omitempty"`

	// 检测详情
	Details *[]TextDetectionResultDetail `json:"details,omitempty"`
}

func (o TextDetectionResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TextDetectionResult struct{}"
	}

	return strings.Join([]string{"TextDetectionResult", string(data)}, " ")
}
