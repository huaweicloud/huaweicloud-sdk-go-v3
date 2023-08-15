package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateCertificateByCsrRequestBody struct {

	// 父CA证书ID。
	IssuerId string `json:"issuer_id"`

	// 证书签名请求。请使用“\\r\\n”或“\\n”替代证书签名请求中的换行符，若通过console端请求此接口，则无需做符号转换。
	Csr string `json:"csr"`

	Validity *Validity `json:"validity"`

	// 证书类型，用于区分从属CA与私有证书。   - **ENTITY_CERT** : 签发私有证书，为缺省值；   - **INTERMEDIATE_CA** : 签发从属CA。
	Type *string `json:"type,omitempty"`

	// 路径长度，仅当签发从属CA时有效。
	PathLength *int32 `json:"path_length,omitempty"`

	// 主体备用名称(本接口预留参数，当前在后端被忽略)，详情请参见**SubjectAlternativeName**字段数据结构说明。
	SubjectAlternativeNames *[]SubjectAlternativeName `json:"subject_alternative_names,omitempty"`
}

func (o CreateCertificateByCsrRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCertificateByCsrRequestBody struct{}"
	}

	return strings.Join([]string{"CreateCertificateByCsrRequestBody", string(data)}, " ")
}
