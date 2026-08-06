package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListKeyPolicyResponseBodyKeyspacePolicies struct {

	// **参数解释：** 密钥策略ID **取值范围：** 不涉及
	PolicyId string `json:"policy_id"`

	// **参数解释：** 密钥策略名称 **取值范围：** 不涉及
	PolicyName string `json:"policy_name"`

	// **参数解释：** 密钥空间ID **取值范围：** 不涉及
	KeyspaceId string `json:"keyspace_id"`

	Policy *ListKeyPolicyResponseBodyPolicy `json:"policy"`

	// **参数解释：** 密钥策略描述信息 **取值范围：** 不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释：** 密钥策略创建人 **取值范围：** 不涉及
	CreatedBy string `json:"created_by"`

	// **参数解释：** 密钥策略创建时间 **取值范围：** 不涉及
	CreateTime string `json:"create_time"`

	// **参数解释：** 密钥策略最近修改时间 **取值范围：** 不涉及
	LastModifyTime string `json:"last_modify_time"`

	// **参数解释：** 密钥策略最近访问时间 **取值范围：** 不涉及
	LastAccessTime string `json:"last_access_time"`
}

func (o ListKeyPolicyResponseBodyKeyspacePolicies) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListKeyPolicyResponseBodyKeyspacePolicies struct{}"
	}

	return strings.Join([]string{"ListKeyPolicyResponseBodyKeyspacePolicies", string(data)}, " ")
}
