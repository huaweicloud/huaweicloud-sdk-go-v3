package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SyncNodeResponse Response Object
type SyncNodeResponse struct {

	// **参数解释**： 固定值\"Sync node success\"，表示同步节点成功。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SyncNodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SyncNodeResponse struct{}"
	}

	return strings.Join([]string{"SyncNodeResponse", string(data)}, " ")
}
