package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListCertificatesRequest struct {
	// 每页条目数量，取值如下： - 10：每页显示10条证书信息。 - 20：每页显示20条证书信息。 - 50：每页显示50条证书信息。

	Limit *int32 `json:"limit,omitempty"`
	// 偏移量。取值范围为1~30。

	Offset *int32 `json:"offset,omitempty"`
	// 排序方式。根据排序参数sort_key进行排序，取值如下： - ASC：升序。 - DESC：降序。

	SortDir *string `json:"sort_dir,omitempty"`
	// 排序依据参数，取值如下： - certExpiredTime：证书到期时间。 - certStatus：证书状态。 - certUpdateTime：证书更新时间。

	SortKey *string `json:"sort_key,omitempty"`
	// 证书状态，取值如下： - ALL：所有证书状态。 - PAID：证书已支付，待申请证书。 - ISSUED：证书已签发。 - CHECKING：证书申请审核中。 - CANCELCHECKING：取消证书申请审核中。 - UNPASSED：证书申请未通过。 - EXPIRED：证书已过期。 - REVOKING：证书吊销申请审核中。 - REVOKED：证书已吊销。 - UPLOAD：证书托管中。 - CHECKING_ORG：待完成企业资格认证。 - ISSUING：证书待签发。 - SUPPLEMENTCHECKING：多域名证书新增附加域名审核中。

	Status *string `json:"status,omitempty"`
}

func (o ListCertificatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCertificatesRequest struct{}"
	}

	return strings.Join([]string{"ListCertificatesRequest", string(data)}, " ")
}
