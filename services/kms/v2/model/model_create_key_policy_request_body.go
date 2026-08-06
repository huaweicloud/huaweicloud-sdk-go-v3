package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateKeyPolicyRequestBody struct {

	// **参数解释：** 密钥策略归属的可信密钥空间ID **约束限制：** 满足正则表达式^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$ **取值范围：** 不涉及 **默认取值：** 不涉及
	KeyspaceId string `json:"keyspace_id"`

	// **参数解释：** 策略策略名称 **约束限制：** 满足正则表达式^[a-zA-Z0-9:/_-]{1,255}$ **取值范围：** 1-255 **默认取值：** 不涉及
	PolicyName string `json:"policy_name"`

	// **参数解释：** 密钥策略 **约束限制：** 转移后的JSON字符串 **取值范围：** 不涉及 **默认取值：** 不涉及
	Policy string `json:"policy"`

	// **参数解释：** 密钥策略描述信息 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Description *string `json:"description,omitempty"`
}

func (o CreateKeyPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateKeyPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"CreateKeyPolicyRequestBody", string(data)}, " ")
}
