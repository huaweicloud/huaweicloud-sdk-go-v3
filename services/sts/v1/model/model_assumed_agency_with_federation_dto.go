package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssumedAgencyWithFederationDto **参数解释**： 信任委托会话信息。  **取值范围**： 不涉及。
type AssumedAgencyWithFederationDto struct {

	// **参数解释**： 信任委托会话的URN。  **取值范围**： 不涉及。
	Urn string `json:"urn"`

	// **参数解释**： 信任委托会话的唯一标识，包含了信任委托ID和信任委托会话名称信息。  **取值范围**： 不涉及。
	Id string `json:"id"`
}

func (o AssumedAgencyWithFederationDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssumedAgencyWithFederationDto struct{}"
	}

	return strings.Join([]string{"AssumedAgencyWithFederationDto", string(data)}, " ")
}
