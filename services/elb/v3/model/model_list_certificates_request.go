package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCertificatesRequest Request Object
type ListCertificatesRequest struct {

	// **参数解释**：上一页最后一条记录的ID。  **约束限制**： - 必须与limit一起使用。 - 不指定时表示查询第一页。 - 该字段不允许为空或无效的ID。  **取值范围**：不涉及  **默认取值**：不涉及
	Marker *string `json:"marker,omitempty"`

	// **参数解释**：每页返回的个数。  **约束限制**：不涉及  **取值范围**：0-2000  **默认取值**：2000
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**：是否反向查询。  **约束限制**： - 必须与limit一起使用。 - 当page_reverse=true时，若要查询上一页，marker取值为当前页返回值的previous_marker。  **取值范围**： - true：查询上一页。 - false：查询下一页。  **默认取值**：false
	PageReverse *bool `json:"page_reverse,omitempty"`

	// **参数解释**：证书ID。 支持多值查询，查询条件格式：*id=xxx&id=xxx*。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	Id *[]string `json:"id,omitempty"`

	// **参数解释**：证书的名称。 支持多值查询，查询条件格式：*name=xxx&name=xxx*。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	Name *[]string `json:"name,omitempty"`

	// **参数解释**：证书的描述。 支持多值查询，查询条件格式：*description=xxx&description=xxx*。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	Description *[]string `json:"description,omitempty"`

	// **参数解释**：证书的管理状态。  **约束限制**：不涉及  **取值范围**： - true：表示证书可用。 - false：表示证书不可用。  **默认取值**：不涉及
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// **参数解释**：服务器证书所签域名。 支持多值查询，查询条件格式：domain=xxx&domain=xxx。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	Domain *[]string `json:"domain,omitempty"`

	// **参数解释**：证书的类型。 支持多值查询，查询条件格式：type=xxx&type=xxx。  **约束限制**：不涉及  **取值范围**： - server：服务器证书。 - client：CA证书。 - server_sm：服务器SM双证书。  **默认取值**：不涉及
	Type *[]string `json:"type,omitempty"`

	// **参数解释**：云证书管理服务（CCM）中的证书ID。 支持多值查询，查询条件格式：scm_certificate_id=xxx&scm_certificate_id=xxx。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	ScmCertificateId *[]string `json:"scm_certificate_id,omitempty"`

	// **参数解释**：证书的主域名。 支持多值查询，查询条件格式：common_name=xxx&common_name=xxx。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	CommonName *[]string `json:"common_name,omitempty"`

	// **参数解释**：证书的指纹。 支持多值查询，查询条件格式：fingerprint=xxx&fingerprint=xxx。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	Fingerprint *[]string `json:"fingerprint,omitempty"`

	// **参数解释**：证书来源。 支持多值查询，查询条件格式：source=xxx&source=xxx。  **约束限制**：不涉及  **取值范围**： - scm：表示关联云证书管理服务（CCM）中的证书。 - 空值：表示自有证书。  **默认取值**：不涉及
	Source *[]string `json:"source,omitempty"`

	// **参数解释**：修改保护状态。 支持多值查询，查询条件格式：protection_status=xxx&protection_status=xxx。  **约束限制**：不涉及  **取值范围**： - nonProtection: 不保护 - consoleProtection: 控制台修改保护，即禁止通过控制台修改。  **默认取值**：不涉及
	ProtectionStatus *[]string `json:"protection_status,omitempty"`

	// **参数解释**：设置修改保护的原因。 支持多值查询，查询条件格式：protection_reason=xxx&protection_reason=xxx。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	ProtectionReason *[]string `json:"protection_reason,omitempty"`

	// **参数解释**：资源所属的企业项目ID。 支持多值查询，查询条件格式： *enterprise_project_id=xxx&enterprise_project_id=xxx*。  **约束限制**： - 如果enterprise_project_id不传值，默认查询所有企业项目下的资源，鉴权按照细粒度权限鉴权，必须在用户组下分配elb:certificates:list权限。 - 如果enterprise_project_id传值，鉴权按照企业项目权限鉴权，分为传入具体eps_id和all_granted_eps两种场景，前者查询指定eps_id的eps下的资源，后者查询的是所有有list权限的eps下的资源。  **取值范围**：不涉及  **默认取值**：不涉及
	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`
}

func (o ListCertificatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCertificatesRequest struct{}"
	}

	return strings.Join([]string{"ListCertificatesRequest", string(data)}, " ")
}
