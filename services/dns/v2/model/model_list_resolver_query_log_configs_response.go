package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResolverQueryLogConfigsResponse Response Object
type ListResolverQueryLogConfigsResponse struct {

	// 解析器访问日志列表。
	ResolverQueryLogConfigs *[]ResolverQueryLogConfig `json:"resolver_query_log_configs,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListResolverQueryLogConfigsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResolverQueryLogConfigsResponse struct{}"
	}

	return strings.Join([]string{"ListResolverQueryLogConfigsResponse", string(data)}, " ")
}
