package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type ImageMediaTaggingItemBody struct {

	// 置信度，将Float型置信度转为String类型返回,取值范围：0-100。
	Confidence *string `json:"confidence,omitempty" xml:"confidence"`

	// 标签的类别。返回的标签类型，包含二十多种大类，具体可以参考[[图像标签](http://support.huaweicloud.com/image_faq/image_01_0037.html)](tag:hc)[[图像标签](https://support.huaweicloud.com/intl/zh-cn/image_faq/image_01_0037.html)](tag:hk)
	Type *string `json:"type,omitempty" xml:"type"`

	// 标签名称。
	Tag *string `json:"tag,omitempty" xml:"tag"`

	I18nTag *ImageMediaTaggingItemBodyI18nTag `json:"i18n_tag,omitempty" xml:"i18n_tag"`

	I18nType *ImageMediaTaggingItemBodyI18nType `json:"i18n_type,omitempty" xml:"i18n_type"`

	// 目标检测框信息，为空则表示没有目标检测框。
	Instances *[]ImageTaggingInstance `json:"instances,omitempty" xml:"instances"`
}

func (o ImageMediaTaggingItemBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageMediaTaggingItemBody struct{}"
	}

	return strings.Join([]string{"ImageMediaTaggingItemBody", string(data)}, " ")
}