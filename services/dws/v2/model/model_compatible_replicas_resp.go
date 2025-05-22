package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CompatibleReplicasResp struct {

	// **参数解释**： ID。 **取值范围**： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 名称。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 链接。 **取值范围**： 不涉及。
	Links *[]LinkResp `json:"links,omitempty"`
}

func (o CompatibleReplicasResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompatibleReplicasResp struct{}"
	}

	return strings.Join([]string{"CompatibleReplicasResp", string(data)}, " ")
}
