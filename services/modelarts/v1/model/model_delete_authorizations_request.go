package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAuthorizationsRequest Request Object
type DeleteAuthorizationsRequest struct {

	// **参数解释**：用户ID，当user_id为all时，表示删除所有IAM子用户的授权。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	UserId *string `json:"user_id,omitempty"`
}

func (o DeleteAuthorizationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAuthorizationsRequest struct{}"
	}

	return strings.Join([]string{"DeleteAuthorizationsRequest", string(data)}, " ")
}
