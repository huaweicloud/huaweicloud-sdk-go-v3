package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckGrowClusterResponse Response Object
type CheckGrowClusterResponse struct {

	// **参数解释**： 扩容前检查信息。 **取值范围**： 不涉及。
	Data *[]GrowCheckResult `json:"data,omitempty"`

	Note           *NoteInfo `json:"note,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o CheckGrowClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckGrowClusterResponse struct{}"
	}

	return strings.Join([]string{"CheckGrowClusterResponse", string(data)}, " ")
}
