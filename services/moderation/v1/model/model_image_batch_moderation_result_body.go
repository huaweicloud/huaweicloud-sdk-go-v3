package model

import (
	"encoding/json"

	"strings"
)

type ImageBatchModerationResultBody struct {
	// 图片的URL路径。

	Url *string `json:"url,omitempty"`
	// 请参见[表2](https://support.huaweicloud.com/api-moderation/moderation_03_0019.html#moderation_03_0019__zh-cn_topic_0087142444_table663805019540)中suggestion字段说明。

	Suggestion *string `json:"suggestion,omitempty"`

	Detail *ImageDetectionResultDetail `json:"detail,omitempty"`
	// 具体每个场景的检测结果。  block：包含敏感信息，不通过  pass：不包含敏感信息，通过  review：需要人工复检

	CategorySuggestions *interface{} `json:"category_suggestions,omitempty"`
}

func (o ImageBatchModerationResultBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ImageBatchModerationResultBody struct{}"
	}

	return strings.Join([]string{"ImageBatchModerationResultBody", string(data)}, " ")
}
