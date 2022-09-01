package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TemplateProductExt struct {

	// 产品id。
	Id *string `json:"id,omitempty" xml:"id"`

	// 产品短名。
	Productshort *string `json:"productshort,omitempty" xml:"productshort"`

	// 产品名。
	ProductName *string `json:"product_name,omitempty" xml:"product_name"`

	// 首页链接。
	HomeLink *string `json:"home_link,omitempty" xml:"home_link"`

	// api调试链接。
	ApiLink *string `json:"api_link,omitempty" xml:"api_link"`

	// sdk下载链接。
	SdkLink *string `json:"sdk_link,omitempty" xml:"sdk_link"`

	// 文档链接。
	DocLink *string `json:"doc_link,omitempty" xml:"doc_link"`

	// logo链接。
	LogoLink *string `json:"logo_link,omitempty" xml:"logo_link"`
}

func (o TemplateProductExt) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateProductExt struct{}"
	}

	return strings.Join([]string{"TemplateProductExt", string(data)}, " ")
}
