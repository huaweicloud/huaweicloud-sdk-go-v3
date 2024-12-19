package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCertificatesRequest Request Object
type ListCertificatesRequest struct {

	// 每页条目数量，取值如下： - 10：每页显示10条证书信息。 - 20：每页显示20条证书信息。 - 50：每页显示50条证书信息。
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量。
	Offset *int32 `json:"offset,omitempty"`

	// 排序方式。根据排序参数sort_key进行排序，取值如下： - ASC：升序。 - DESC：降序。
	SortDir *string `json:"sort_dir,omitempty"`

	// 排序依据参数，取值如下： - certExpiredTime：证书到期时间。 - certStatus：证书状态。 - certUpdateTime：证书更新时间。
	SortKey *string `json:"sort_key,omitempty"`

	// 证书状态，取值如下： - ALL：所有证书状态。 - PAID：证书已支付，待申请证书。 - ISSUED：证书已签发。 - CHECKING：证书申请审核中。 - CANCELCHECKING：取消证书申请审核中。 - UNPASSED：证书申请未通过。 - EXPIRED：证书已过期。 - REVOKING：证书吊销申请审核中。 - REVOKED：证书已吊销。 - UPLOAD：证书托管中。 - CHECKING_ORG：待完成企业资格认证。 - ISSUING：证书待签发。 - SUPPLEMENTCHECKING：多域名证书新增附加域名审核中。
	Status *string `json:"status,omitempty"`

	// 企业多项目ID。用户未开通企业多项目时，不需要输入该字段。 用户开通企业多项目时，查询资源可以输入该字段。 若用户不输入该字段，默认查询租户所有有权限的企业多项目下的资源。 此时“enterprise_project_id”取值为“all”。 若用户输入该字段，取值满足以下任一条件.  取值为“all”  取值为“0”  满足正则匹配：“^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$”
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 是否支持部署。
	DeploySupport *bool `json:"deploy_support,omitempty"`

	// 过滤资源是否属于当前租户，取值如下： - true：只查属于当前租户的资源，不包括共享资源。 - false：查询当前租户及共享给该租户的资源。
	OwnedBySelf *bool `json:"owned_by_self,omitempty"`
}

func (o ListCertificatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCertificatesRequest struct{}"
	}

	return strings.Join([]string{"ListCertificatesRequest", string(data)}, " ")
}
