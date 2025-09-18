package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CopyPolicyByIdRequest Request Object
type CopyPolicyByIdRequest struct {

	// 您可以通过调用企业项目管理服务（EPS)的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 源策略ID
	SrcPolicyId string `json:"src_policy_id"`

	// 复制出的新策略名称
	DestPolicyName string `json:"dest_policy_name"`
}

func (o CopyPolicyByIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyPolicyByIdRequest struct{}"
	}

	return strings.Join([]string{"CopyPolicyByIdRequest", string(data)}, " ")
}
