package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCsrResponse Response Object
type UpdateCsrResponse struct {

	// CSR的ID。
	Id *string `json:"id,omitempty"`

	// CSR名称。
	Name *string `json:"name,omitempty"`

	// CSR内容。
	Csr *string `json:"csr,omitempty"`

	// CSR绑定的域名。
	DomainName *string `json:"domain_name,omitempty"`

	// CSR绑定的附加域名。
	Sans *string `json:"sans,omitempty"`

	// 密钥算法。
	PrivateKeyAlgo *string `json:"private_key_algo,omitempty"`

	// CSR用途。
	Usage *string `json:"usage,omitempty"`

	// 国家。
	CompanyCountry *string `json:"company_country,omitempty"`

	// 省份。
	CompanyProvince *string `json:"company_province,omitempty"`

	// 城市。
	CompanyCity *string `json:"company_city,omitempty"`

	// 公司名称。
	CompanyName *string `json:"company_name,omitempty"`

	// CSR创建时间。
	CreateTime *int64 `json:"create_time,omitempty"`

	// CSR更新时间。
	UpdateTime     *int64 `json:"update_time,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o UpdateCsrResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCsrResponse struct{}"
	}

	return strings.Join([]string{"UpdateCsrResponse", string(data)}, " ")
}
