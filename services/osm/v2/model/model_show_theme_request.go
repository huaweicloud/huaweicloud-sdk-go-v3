package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowThemeRequest Request Object
type ShowThemeRequest struct {

	// 调用智能客服服务标志。
	XServiceKey string `json:"x-service-key"`

	// 站点标记，0-中国站  1-国际站
	XSite *string `json:"x-site,omitempty"`

	// 区域语言简写，en-us  zh-cn
	XLanguage *string `json:"x-language,omitempty"`

	// 产品类型Id
	ProductTypeId *string `json:"product_type_id,omitempty"`

	// 产品类型名称
	ProductTypeName *string `json:"product_type_name,omitempty"`

	// 产品类型缩写
	ProductTypeShortName *string `json:"product_type_short_name,omitempty"`
}

func (o ShowThemeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowThemeRequest struct{}"
	}

	return strings.Join([]string{"ShowThemeRequest", string(data)}, " ")
}
