package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClusterProtectionDefaultPolicyResponse Response Object
type ListClusterProtectionDefaultPolicyResponse struct {

	// **参数解释**: 总数 **取值范围**: 最小值0，最大值10000
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**: 用户Token **取值范围**: 字符长度1-32768位
	XAuthToken *string `json:"x_auth_token,omitempty"`

	// **参数解释**: 项目ID **取值范围**: 字符长度1-32768位
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释**: 主机所属的企业项目ID **取值范围**: 字符长度1-256位
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: Region ID **取值范围**: 字符长度1-32768位
	Region *string `json:"region,omitempty"`

	// **参数解释**: 通用策略数 **取值范围**: 最小值0，最大值2147483647
	GeneralPolicyNum *int32 `json:"general_policy_num,omitempty"`

	// **参数解释**: 恶意镜像策略数 **取值范围**: 最小值0，最大值2147483647
	MaliciousImagePolicyNum *int32 `json:"malicious_image_policy_num,omitempty"`

	// **参数解释**: 安全镜像策略数 **取值范围**: 最小值0，最大值2147483647
	SecurityPolicyNum *int32 `json:"security_policy_num,omitempty"`

	// **参数解释**: 集群防护策略列表 **取值范围**: 取值0-10000个ClusterPolicyResponseInfo对象
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
