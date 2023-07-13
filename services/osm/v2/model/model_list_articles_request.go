package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListArticlesRequest Request Object
type ListArticlesRequest struct {

	// 调用智能客服服务标志。
	XServiceKey *string `json:"x-service-key,omitempty"`

	// 站点标记，0-中国站  1-国际站
	XSite *string `json:"x-site,omitempty"`

	// 区域语言简写，en-us  zh-cn
	XLanguage *string `json:"x-language,omitempty"`

	// 搜索类型：HOT-热点案例,RECOMMEND-推荐案例
	SearchType string `json:"search_type"`

	Body *SearchArticlesReq `json:"body,omitempty"`
}

func (o ListArticlesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListArticlesRequest struct{}"
	}

	return strings.Join([]string{"ListArticlesRequest", string(data)}, " ")
}
