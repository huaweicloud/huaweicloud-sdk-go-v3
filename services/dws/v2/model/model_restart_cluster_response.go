package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestartClusterResponse Response Object
type RestartClusterResponse struct {

	// **参数解释**： 请求成功时的空白响应。 **取值范围**： 不涉及。
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o RestartClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartClusterResponse struct{}"
	}

	return strings.Join([]string{"RestartClusterResponse", string(data)}, " ")
}
