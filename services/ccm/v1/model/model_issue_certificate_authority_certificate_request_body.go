package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

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
}

func (o IssueCertificateAuthorityCertificateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueCertificateAuthorityCertificateRequestBody struct{}"
	}

	return strings.Join([]string{"IssueCertificateAuthorityCertificateRequestBody", string(data)}, " ")
}
