package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RefreshOperationalTaskResponse Response Object
type RefreshOperationalTaskResponse struct {

	// **参数解释**： 请求成功时的空白响应。 **取值范围**： 不涉及。
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o RefreshOperationalTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RefreshOperationalTaskResponse struct{}"
	}

	return strings.Join([]string{"RefreshOperationalTaskResponse", string(data)}, " ")
}
