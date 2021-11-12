package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateDomainV2Response struct {
	// 自定义域名

	UrlDomain *string `json:"url_domain,omitempty"`
	// 自定义域名的编号

	Id *string `json:"id,omitempty"`
	// CNAME解析状态 - 1: 未解析 - 2: 解析中 - 3: 解析成功 - 4: 解析失败

	Status *int32 `json:"status,omitempty"`
	// 支持的最小SSL版本

	MinSslVersion  *string `json:"min_ssl_version,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateDomainV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDomainV2Response struct{}"
	}

	return strings.Join([]string{"UpdateDomainV2Response", string(data)}, " ")
}
