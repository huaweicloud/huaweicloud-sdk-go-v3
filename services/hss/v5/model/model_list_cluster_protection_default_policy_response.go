package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClusterProtectionDefaultPolicyResponse Response Object
type ListClusterProtectionDefaultPolicyResponse struct {

	// 总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 用户Token。
	XAuthToken *string `json:"x_auth_token,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 企业ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// Region
	Region *string `json:"region,omitempty"`

	// general策略数
	GeneralPolicyNum *int32 `json:"general_policy_num,omitempty"`

	// malicious_image策略数
	MaliciousImagePolicyNum *int32 `json:"malicious_image_policy_num,omitempty"`

	// security_policy策略数
	SecurityPolicyNum *int32 `json:"security_policy_num,omitempty"`

	// 集群防护策略列表
	DataList       *[]ClusterPolicyResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ListClusterProtectionDefaultPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClusterProtectionDefaultPolicyResponse struct{}"
	}

	return strings.Join([]string{"ListClusterProtectionDefaultPolicyResponse", string(data)}, " ")
}
