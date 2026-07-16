package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetAuthorizationsResponse Response Object
type GetAuthorizationsResponse struct {

	// **参数解释**：授权信息总数。 **取值范围**：不涉及。
	TotalCount float32 `json:"total_count,omitempty"`

	// **参数解释**：授权信息列表。
	Auth           *[]AuthorizationResponse `json:"auth,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o GetAuthorizationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetAuthorizationsResponse struct{}"
	}

	return strings.Join([]string{"GetAuthorizationsResponse", string(data)}, " ")
}
