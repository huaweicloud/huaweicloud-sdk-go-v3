package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNoticesRequest Request Object
type ListNoticesRequest struct {

	// 调用智能客服服务标志。
	XServiceKey *string `json:"x-service-key,omitempty"`

	// 站点标记，0-中国站  1-国际站
	XSite *string `json:"x-site,omitempty"`

	// 区域语言简写，en-us  zh-cn
	XLanguage *string `json:"x-language,omitempty"`

	Body *SearchNoticesReq `json:"body,omitempty"`
}

func (o ListNoticesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNoticesRequest struct{}"
	}

	return strings.Join([]string{"ListNoticesRequest", string(data)}, " ")
}
