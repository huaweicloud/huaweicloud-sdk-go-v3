package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UrlRedirectionOptions url重定向配置。
type UrlRedirectionOptions struct {

	// 允许重定向url。
	UrlIncludeList *string `json:"url_include_list,omitempty"`

	// 拒绝重定向url。
	UrlExcludeList *string `json:"url_exclude_list,omitempty"`
}

func (o UrlRedirectionOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UrlRedirectionOptions struct{}"
	}

	return strings.Join([]string{"UrlRedirectionOptions", string(data)}, " ")
}
