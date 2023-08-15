package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RewriteUrlConfig 转发到的后端主机组的URL配置。rewrite_url_enable为true时，改字段有效。
type RewriteUrlConfig struct {

	// url host
	Host *string `json:"host,omitempty"`

	// url路径
	Path *string `json:"path,omitempty"`

	// url查询字符串
	Query *string `json:"query,omitempty"`
}

func (o RewriteUrlConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RewriteUrlConfig struct{}"
	}

	return strings.Join([]string{"RewriteUrlConfig", string(data)}, " ")
}
