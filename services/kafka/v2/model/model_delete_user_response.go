package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteUserResponse Response Object
type DeleteUserResponse struct {

	// **参数解释**： 响应信息。 **取值范围**： 不涉及。
	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteUserResponse struct{}"
	}

	return strings.Join([]string{"DeleteUserResponse", string(data)}, " ")
}
