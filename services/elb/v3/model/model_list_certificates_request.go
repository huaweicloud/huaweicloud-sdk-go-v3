package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListCertificatesRequest struct {
	// SSL证书的管理状态；暂不支持。

	AdminStateUp *bool `json:"admin_state_up,omitempty"`
	// SSL证书的描述。

	Description *[]string `json:"description,omitempty"`
	// 服务器证书所签域名。该字段仅type为server时有效。

	Domain *[]string `json:"domain,omitempty"`
	// 证书ID。

	Id *[]string `json:"id,omitempty"`
	// 每页返回的个数。

	Limit *int32 `json:"limit,omitempty"`
	// 上一页最后一条记录的ID。  使用说明：  - 必须与limit一起使用。 - 不指定时表示查询第一页。 - 该字段不允许为空或无效的ID。

	Marker *string `json:"marker,omitempty"`
	// SSL证书的名称。

	Name *[]string `json:"name,omitempty"`
	// 分页的顺序，true表示从后往前分页，false表示从前往后分页，默认为false。使用说明：必须与limit一起使用。

	PageReverse *bool `json:"page_reverse,omitempty"`
	// SSL证书的类型。分为服务器证书(server)和CA证书(client)。

	Type *[]string `json:"type,omitempty"`
}

func (o ListCertificatesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListCertificatesRequest struct{}"
	}

	return strings.Join([]string{"ListCertificatesRequest", string(data)}, " ")
}
