package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssumedAgencyDto **参数解释**： 委托会话或信任委托会话信息。  **取值范围**： 不涉及。
type AssumedAgencyDto struct {

	// **参数解释**： 委托会话或信任委托会话的URN。  **取值范围**： 不涉及。
	Urn string `json:"urn"`

	// **参数解释**： 委托会话或信任委托会话的唯一标识，包含了委托ID和委托会话名称信息。  **取值范围**： 不涉及。
	Id string `json:"id"`
}

func (o AssumedAgencyDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssumedAgencyDto struct{}"
	}

	return strings.Join([]string{"AssumedAgencyDto", string(data)}, " ")
}
