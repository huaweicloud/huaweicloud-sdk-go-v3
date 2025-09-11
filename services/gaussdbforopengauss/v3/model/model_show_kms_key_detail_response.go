package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowKmsKeyDetailResponse Response Object
type ShowKmsKeyDetailResponse struct {

	// **参数解释**: 秘钥ID。 **取值范围**: 不涉及。
	KeyId *string `json:"key_id,omitempty"`

	// **参数解释**: 密钥别名。 **取值范围**: 不涉及。
	KeyAlias *string `json:"key_alias,omitempty"`

	// **参数解释**: 用户域ID。 **取值范围**: 不涉及。
	DomainId *string `json:"domain_id,omitempty"`

	// **参数解释**: 秘钥状态。 **取值范围**: - “1”表示待激活状态。 - “2”表示启用状态。 - “3”表示禁用状态。 - “4”表示计划删除状态。 - “5”表示等待导入状态。
	KeyState       *string `json:"key_state,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowKmsKeyDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKmsKeyDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowKmsKeyDetailResponse", string(data)}, " ")
}
