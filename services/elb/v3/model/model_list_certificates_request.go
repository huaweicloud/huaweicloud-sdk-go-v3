package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListCertificatesRequest struct {
	// 上一页最后一条记录的ID。  使用说明： - 必须与limit一起使用。 - 不指定时表示查询第一页。 - 该字段不允许为空或无效的ID。

	Marker *string `json:"marker,omitempty"`
	// 每页返回的个数。

	Limit *int32 `json:"limit,omitempty"`
	// 分页的顺序，true表示从后往前分页，false表示从前往后分页，默认为false。  使用说明： - 必须与limit一起使用。

	PageReverse *bool `json:"page_reverse,omitempty"`
	// 证书ID。  支持多值查询，查询条件格式：*id=xxx&id=xxx*。

	Id *[]string `json:"id,omitempty"`
	// 证书的名称。  支持多值查询，查询条件格式：*name=xxx&name=xxx*。

	Name *[]string `json:"name,omitempty"`
	// 证书的描述。  支持多值查询，查询条件格式：*description=xxx&description=xxx*。

	Description *[]string `json:"description,omitempty"`
	// 证书的管理状态。  不支持该字段，请勿使用。

	AdminStateUp *bool `json:"admin_state_up,omitempty"`
	// 服务器证书所签域名。该字段仅type为server时有效。  支持多值查询，查询条件格式：domain=xxx&domain=xxx。

	Domain *[]string `json:"domain,omitempty"`
	// 证书的类型。分为服务器证书(server)和CA证书(client)。  支持多值查询，查询条件格式：type=xxx&type=xxx。

	Type *[]string `json:"type,omitempty"`
}

func (o ListCertificatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCertificatesRequest struct{}"
	}

	return strings.Join([]string{"ListCertificatesRequest", string(data)}, " ")
}
