package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListToolsRequest Request Object
type ListToolsRequest struct {

	// 调用智能客服服务标志。
	XServiceKey *string `json:"x-service-key,omitempty"`

	// 站点标记，0-中国站  1-国际站
	XSite *string `json:"x-site,omitempty"`

	// 区域语言简写，en-us  zh-cn
	XLanguage *string `json:"x-language,omitempty"`

	Body *SearchToolsReq `json:"body,omitempty"`
}

func (o ListToolsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListToolsRequest struct{}"
	}

	return strings.Join([]string{"ListToolsRequest", string(data)}, " ")
}
