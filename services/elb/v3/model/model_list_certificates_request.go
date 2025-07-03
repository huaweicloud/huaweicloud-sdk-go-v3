package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCertificatesRequest Request Object
type ListCertificatesRequest struct {

	// 上一页最后一条记录的ID。  使用说明： - 必须与limit一起使用。 - 不指定时表示查询第一页。 - 该字段不允许为空或无效的ID。
	Marker *string `json:"marker,omitempty"`

	// 参数解释：每页返回的个数。  取值范围：0-2000  默认取值：2000
	Limit *int32 `json:"limit,omitempty"`

	// 是否反向查询。  取值： - true：查询上一页。 - false：查询下一页，默认。  使用说明： - 必须与limit一起使用。 - 当page_reverse=true时，若要查询上一页，marker取值为当前页返回值的previous_marker。
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

	// SCM证书ID。  支持多值查询，查询条件格式：scm_certificate_id=xxx&scm_certificate_id=xxx。
	ScmCertificateId *[]string `json:"scm_certificate_id,omitempty"`

	// 证书的主域名。  支持多值查询，查询条件格式：common_name=xxx&common_name=xxx。
	CommonName *[]string `json:"common_name,omitempty"`

	// 证书的指纹。  支持多值查询，查询条件格式：fingerprint=xxx&fingerprint=xxx。
	Fingerprint *[]string `json:"fingerprint,omitempty"`

	// 证书来源。  支持多值查询，查询条件格式：source=xxx&source=xxx。
	Source *[]string `json:"source,omitempty"`

	// 修改保护状态。  支持多值查询，查询条件格式：protection_status=xxx&protection_status=xxx。
	ProtectionStatus *[]string `json:"protection_status,omitempty"`

	// 设置修改保护的原因。  支持多值查询，查询条件格式：protection_reason=xxx&protection_reason=xxx。
	ProtectionReason *[]string `json:"protection_reason,omitempty"`

	// 参数解释：所属的企业项目ID。 如果enterprise_project_id不传值，默认查询所有企业项目下的资源，鉴权按照细粒度权限鉴权，必须在用户组下分配elb:certificates:list权限。 如果enterprise_project_id传值，鉴权按照企业项目权限鉴权，分为传入具体eps_id和all_granted_eps两种场景，前者查询指定eps_id的eps下的资源，后者查询的是所有有list权限的eps下的资源。  支持多值查询，查询条件格式： *enterprise_project_id=xxx&enterprise_project_id=xxx*。  [不支持该字段，请勿使用。](tag:dt,hcso_dt)
	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`
}

func (o ListCertificatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCertificatesRequest struct{}"
	}

	return strings.Join([]string{"ListCertificatesRequest", string(data)}, " ")
}
