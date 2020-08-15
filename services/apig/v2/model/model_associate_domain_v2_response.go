/*
 * APIG
 *
 * API网关（API Gateway）是为开发者、合作伙伴提供的高性能、高可用、高安全的API托管服务，帮助用户轻松构建、管理和发布任意规模的API。
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

// Response Object
type AssociateDomainV2Response struct {
	// 自定义域名
	UrlDomain string `json:"url_domain,omitempty"`
	// 域名的编号
	Id string `json:"id,omitempty"`
	// CNAME解析状态 - 1: 未解析 - 2: 解析中 - 3: 解析成功 - 4: 解析失败
	Status int32 `json:"status,omitempty"`
}

func (o AssociateDomainV2Response) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"AssociateDomainV2Response", string(data)}, " ")
}
