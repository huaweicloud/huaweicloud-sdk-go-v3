package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IssueCertificateAuthorityCertificateRequestBody struct {

	// 父CA证书ID。
	IssuerId string `json:"issuer_id"`

	// 路径长度。
	PathLength int32 `json:"path_length"`

	// 签名哈希算法，可选值如下：   - **SHA256**   - **SHA384**   - **SHA512**   - **SM3**（中国站）
	SignatureAlgorithm string `json:"signature_algorithm"`

	Validity *Validity `json:"validity"`

	// 企业多项目ID。用户未开通企业多项目时，不需要输入该字段。 用户开通企业多项目时，查询资源可以输入该字段。 若用户不输入该字段，默认查询租户所有有权限的企业多项目下的资源。 此时“enterprise_project_id”取值为“all”。 若用户输入该字段，取值满足以下任一条件.   取值为“all”   取值为“0”   满足正则匹配：“^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$”
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o IssueCertificateAuthorityCertificateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueCertificateAuthorityCertificateRequestBody struct{}"
	}

	return strings.Join([]string{"IssueCertificateAuthorityCertificateRequestBody", string(data)}, " ")
}
