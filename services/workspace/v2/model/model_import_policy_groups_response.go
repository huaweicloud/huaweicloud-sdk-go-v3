package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportPolicyGroupsResponse Response Object
type ImportPolicyGroupsResponse struct {

	// 错误码。
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误描述。
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 加密后的详细拒绝原因，用户可以自行调用STS服务的decode-authorization-message接口进行解密。
	EncodedAuthorizationMessage *string `json:"encoded_authorization_message,omitempty"`

	// 导入策略组所有名字列表。
	PolicyGroupNameList *[]PolicyGroupNameInfo `json:"policy_group_name_list,omitempty"`

	// 导入策略组失败名字列表。
	FailedPolicyGroupNameList *[]PolicyGroupNameInfo `json:"failed_policy_group_name_list,omitempty"`
	HttpStatusCode            int                    `json:"-"`
}

func (o ImportPolicyGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportPolicyGroupsResponse struct{}"
	}

	return strings.Join([]string{"ImportPolicyGroupsResponse", string(data)}, " ")
}
