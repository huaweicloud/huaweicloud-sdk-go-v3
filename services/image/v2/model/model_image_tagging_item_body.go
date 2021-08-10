package model

import (
	"encoding/json"

	"strings"
)

//
type ImageTaggingItemBody struct {
	// 置信度，将Float型置信度转为String类型返回,取值范围：0-100。

	Confidence *string `json:"confidence,omitempty"`
	// 标签的类别。有以下三种类别：  object：实体标签  scene：场景标签  concept：概念标签

	Type *string `json:"type,omitempty"`
	// 标签名称。

	Tag *string `json:"tag,omitempty"`

	I18nTag *ImageTaggingItemBodyI18nTag `json:"i18n_tag,omitempty"`

	I18nType *ImageTaggingItemBodyI18nType `json:"i18n_type,omitempty"`
	// 目标检测框信息，为空则表示没有目标检测框。

	Instances *[]ImageTaggingInstance `json:"instances,omitempty"`
}

func (o ImageTaggingItemBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ImageTaggingItemBody struct{}"
	}

	return strings.Join([]string{"ImageTaggingItemBody", string(data)}, " ")
}
