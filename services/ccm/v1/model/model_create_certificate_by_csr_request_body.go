package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

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

	// 企业多项目ID。用户未开通企业多项目时，不需要输入该字段。 用户开通企业多项目时，查询资源可以输入该字段。 若用户不输入该字段，默认查询租户所有有权限的企业多项目下的资源。 此时“enterprise_project_id”取值为“all”。 若用户输入该字段，取值满足以下任一条件.   取值为“all”   取值为“0”   满足正则匹配：“^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$”
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o CreateCertificateByCsrRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCertificateByCsrRequestBody struct{}"
	}

	return strings.Join([]string{"CreateCertificateByCsrRequestBody", string(data)}, " ")
}
