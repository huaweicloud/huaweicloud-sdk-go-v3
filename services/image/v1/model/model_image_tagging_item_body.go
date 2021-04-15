package model

import (
	"encoding/json"

	"strings"
)

//
type ImageTaggingItemBody struct {
	// 置信度，取值范围：0-100。

	Confidence *string `json:"confidence,omitempty"`
	// 标签的类别 object：实体标签 scene：场景标签 concept：概念标签

	Type *string `json:"type,omitempty"`
	// 标签名称。

	Tag *string `json:"tag,omitempty"`

	I18nTag *ImageTaggingI18nTag `json:"i18n_tag,omitempty"`
}

func (o ImageTaggingItemBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ImageTaggingItemBody struct{}"
	}

	return strings.Join([]string{"ImageTaggingItemBody", string(data)}, " ")
}
