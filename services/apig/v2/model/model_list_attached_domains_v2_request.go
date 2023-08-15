package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAttachedDomainsV2Request Request Object
type ListAttachedDomainsV2Request struct {

	// 证书的编号
	CertificateId string `json:"certificate_id"`

	// 偏移量，表示从此偏移量开始查询，偏移量小于0时，自动转换为0
	Offset *int64 `json:"offset,omitempty"`

	// 每页显示的条目数量，条目数量小于等于0时，自动转换为20，条目数量大于500时，自动转换为500
	Limit *int32 `json:"limit,omitempty"`

	// 独立域名
	UrlDomain *string `json:"url_domain,omitempty"`
}

func (o ListAttachedDomainsV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAttachedDomainsV2Request struct{}"
	}

	return strings.Join([]string{"ListAttachedDomainsV2Request", string(data)}, " ")
}
