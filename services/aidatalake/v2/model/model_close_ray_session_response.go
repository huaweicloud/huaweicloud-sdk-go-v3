package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CloseRaySessionResponse Response Object
type CloseRaySessionResponse struct {

	// **参数解释**：会话ID。 **约束限制**：不涉及。 **取值范围**：长度为1~36的英文、数字、中划线的组合。 **默认取值**：不涉及。
	SessionId *string `json:"session_id,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CloseRaySessionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloseRaySessionResponse struct{}"
	}

	return strings.Join([]string{"CloseRaySessionResponse", string(data)}, " ")
}
