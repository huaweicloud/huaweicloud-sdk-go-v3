package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListKeyPolicyResponse Response Object
type ListKeyPolicyResponse struct {
	PageInfo *ListAccessPointResponseBodyPageInfo `json:"page_info,omitempty"`

	// **参数解释：** 密钥策略列表 **取值范围：** 不涉及
	KeyspacePolicies *[]ListKeyPolicyResponseBodyKeyspacePolicies `json:"keyspace_policies,omitempty"`
	HttpStatusCode   int                                          `json:"-"`
}

func (o ListKeyPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListKeyPolicyResponse struct{}"
	}

	return strings.Join([]string{"ListKeyPolicyResponse", string(data)}, " ")
}
