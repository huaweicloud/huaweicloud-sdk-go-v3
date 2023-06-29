package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCertificateRequest Request Object
type ListCertificateRequest struct {

	// 指定查询返回记录条数，默认值10。
	Limit *int32 `json:"limit,omitempty"`

	// 私有证书名称，返回名称带有name字段的证书集合。
	Name *string `json:"name,omitempty"`

	// 索引位置，从offset指定的下一条数据开始查询。
	Offset *int32 `json:"offset,omitempty"`

	// 私有证书状态，通过状态过滤证书集合。   - **ISSUED** : 已签发；   - **REVOKED** : 已吊销；   - **EXPIRED** : 已过期。
	Status *string `json:"status,omitempty"`

	// 排序属性，目前支持以下属性： - **create_time** : 证书创建时间（默认） - **common_name** : 证书名称 - **issuer_name** : 签发CA名称 - **not_after** : 证书到期时间
	SortKey *string `json:"sort_key,omitempty"`

	// 排序方向，支持以下值：   - **DESC** : 降序（默认）   - **ASC** : 升序
	SortDir *string `json:"sort_dir,omitempty"`
}

func (o ListCertificateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCertificateRequest struct{}"
	}

	return strings.Join([]string{"ListCertificateRequest", string(data)}, " ")
}
