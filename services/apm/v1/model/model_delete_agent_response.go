package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteAgentResponse struct {

	// 删除状态。
	DeleteStatus   *int32 `json:"delete_status,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o DeleteAgentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAgentResponse struct{}"
	}

	return strings.Join([]string{"DeleteAgentResponse", string(data)}, " ")
}
