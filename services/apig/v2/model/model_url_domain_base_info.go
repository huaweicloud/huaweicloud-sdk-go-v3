package model

import (
	"encoding/json"

	"strings"
)

type UrlDomainBaseInfo struct {
	// 自定义域名

	UrlDomain string `json:"url_domain"`
	// 自定义域名的编号

	Id string `json:"id"`
	// CNAME解析状态 - 1: 未解析 - 2: 解析中 - 3: 解析成功 - 4: 解析失败

	Status int32 `json:"status"`
	// 支持的最小SSL版本

	MinSslVersion string `json:"min_ssl_version"`
}

func (o UrlDomainBaseInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UrlDomainBaseInfo struct{}"
	}

	return strings.Join([]string{"UrlDomainBaseInfo", string(data)}, " ")
}
