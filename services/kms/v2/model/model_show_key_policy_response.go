package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowKeyPolicyResponse Response Object
type ShowKeyPolicyResponse struct {

	// **参数解释：** 密钥策略ID **取值范围：** 不涉及
	PolicyId *string `json:"policy_id,omitempty"`

	// **参数解释：** 密钥策略名称 **取值范围：** 不涉及
	PolicyName *string `json:"policy_name,omitempty"`

	// **参数解释：** 密钥空间ID **取值范围：** 不涉及
	KeyspaceId *string `json:"keyspace_id,omitempty"`

	Policy *ShowKeyPolicyResponseBodyPolicy `json:"policy,omitempty"`

	// **参数解释：** 密钥策略描述信息 **取值范围：** 不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释：** 密钥策略创建人 **取值范围：** 不涉及
	CreatedBy *string `json:"created_by,omitempty"`

	// **参数解释：** 密钥策略创建时间 **取值范围：** 不涉及
	CreateTime *string `json:"create_time,omitempty"`

	// **参数解释：** 密钥策略最近修改时间 **取值范围：** 不涉及
	LastModifyTime *string `json:"last_modify_time,omitempty"`

	// **参数解释：** 密钥策略最近访问时间 **取值范围：** 不涉及
	LastAccessTime *string `json:"last_access_time,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowKeyPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKeyPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowKeyPolicyResponse", string(data)}, " ")
}
