package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CsrResponseBody struct {

	// CSR的ID。
	Id string `json:"id"`

	// CSR名称。
	Name string `json:"name"`

	// CSR绑定的域名。
	DomainName string `json:"domain_name"`

	// CSR绑定的附加域名。
	Sans string `json:"sans"`

	// 密钥算法。
	PrivateKeyAlgo string `json:"private_key_algo"`

	// CSR用途。
	Usage string `json:"usage"`

	// 国家。
	CompanyCountry string `json:"company_country"`

	// 省份。
	CompanyProvince string `json:"company_province"`

	// 城市。
	CompanyCity string `json:"company_city"`

	// 公司名称。
	CompanyName string `json:"company_name"`

	// CSR创建时间。
	CreateTime float32 `json:"create_time"`

	// CSR更新时间。
	UpdateTime float32 `json:"update_time"`
}

func (o CsrResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CsrResponseBody struct{}"
	}

	return strings.Join([]string{"CsrResponseBody", string(data)}, " ")
}
