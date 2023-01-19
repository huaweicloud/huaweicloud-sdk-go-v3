package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListCertificateAuthorityRequest struct {

	// 指定查询返回记录条数，默认值10。
	Limit *int32 `json:"limit,omitempty"`

	// CA证书名称（CN）过滤值，用于获取名称中带有特定值的CA证书集合。
	Name *string `json:"name,omitempty"`

	// 索引位置，从offset指定的下一条数据开始查询。默认值为0。
	Offset *int32 `json:"offset,omitempty"`

	// CA证书状态，通过状态过滤证书集合： - **PENDING** : 待激活，此状态下，不可用于签发证书； - **ACTIVED** : 已激活，此状态下，可用于签发证书； - **DISABLED** : 已禁用，此状态下，不可用于签发证书； - **DELETED** : 计划删除，此状态下，不可用于签发证书； - **EXPIRED** : 已过期，此状态下，不可用于签发证书。
	Status *string `json:"status,omitempty"`

	// CA证书类型： - **ROOT** : 根CA证书 - **SUBORDINATE** : 从属CA证书
	Type *string `json:"type,omitempty"`

	// 排序属性，目前支持以下属性： - **create_time** : 证书创建时间（默认） - **common_name** : 证书名称 - **type** : CA证书类型 - **not_after** : 证书到期时间
	SortKey *string `json:"sort_key,omitempty"`

	// 排序方向，支持以下值： - **DESC** : 降序（默认） - **ASC** : 升序
	SortDir *string `json:"sort_dir,omitempty"`
}

func (o ListCertificateAuthorityRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCertificateAuthorityRequest struct{}"
	}

	return strings.Join([]string{"ListCertificateAuthorityRequest", string(data)}, " ")
}
