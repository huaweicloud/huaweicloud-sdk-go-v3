package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDatabaseUserResponse Response Object
type DeleteDatabaseUserResponse struct {

	// **参数解释**： 请求成功时的空白响应。 **取值范围**： 不涉及。
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o DeleteDatabaseUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDatabaseUserResponse struct{}"
	}

	return strings.Join([]string{"DeleteDatabaseUserResponse", string(data)}, " ")
}
