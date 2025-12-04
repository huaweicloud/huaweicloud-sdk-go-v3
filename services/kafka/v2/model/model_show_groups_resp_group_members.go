package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowGroupsRespGroupMembers struct {

	// **参数解释**： 消费组consumer地址。 **取值范围**： 不涉及。
	Host *string `json:"host,omitempty"`

	// **参数解释**： consumer分配到的分区信息。
	Assignment *[]ShowGroupsRespGroupAssignment `json:"assignment,omitempty"`

	// **参数解释**： 消费组consumer的ID。 **取值范围**： 不涉及。
	MemberId *string `json:"member_id,omitempty"`

	// **参数解释**： 客户端ID。 **取值范围**： 不涉及。
	ClientId *string `json:"client_id,omitempty"`
}

func (o ShowGroupsRespGroupMembers) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGroupsRespGroupMembers struct{}"
	}

	return strings.Join([]string{"ShowGroupsRespGroupMembers", string(data)}, " ")
}
