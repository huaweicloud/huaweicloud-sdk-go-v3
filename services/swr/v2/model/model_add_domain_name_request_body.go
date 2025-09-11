package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddDomainNameRequestBody struct {

	// 域名，只能由字母、数字、中划线、星号组成， 星号只能在开头，中划线不能在开头或末 尾，至少包含两个字符串，单个字符串不 超过63个字符，字符串间以点分割，且总 长度不超过100个字符。例如： example.com 或 *.example.com。
	DomainName string `json:"domain_name"`

	// SCM服务的证书ID
	CertificateId string `json:"certificate_id"`
}

func (o AddDomainNameRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddDomainNameRequestBody struct{}"
	}

	return strings.Join([]string{"AddDomainNameRequestBody", string(data)}, " ")
}
