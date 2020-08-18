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
type AssociateCertificateV2Response struct {
	// 证书的名称
	SslName *string `json:"ssl_name,omitempty"`
	// 自定义域名
	UrlDomain *string `json:"url_domain,omitempty"`
	// 证书的编号
	SslId *string `json:"ssl_id,omitempty"`
	// 自定义域名的编号
	Id *string `json:"id,omitempty"`
	// 解析状态值
	Status *int32 `json:"status,omitempty"`
}

func (o AssociateCertificateV2Response) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"AssociateCertificateV2Response", string(data)}, " ")
}
