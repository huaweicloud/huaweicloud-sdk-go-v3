package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckClusterShrinkResponse Response Object
type CheckClusterShrinkResponse struct {

	// **参数解释**： 检查是否通过。如不通过，需要调整缩容节点数重试，或者是当前集群就不满足缩容前置条件。 **取值范围**： 不涉及。
	IsPassed       *bool `json:"is_passed,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o CheckClusterShrinkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckClusterShrinkResponse struct{}"
	}

	return strings.Join([]string{"CheckClusterShrinkResponse", string(data)}, " ")
}
