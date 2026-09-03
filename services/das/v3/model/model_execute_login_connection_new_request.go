package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteLoginConnectionNewRequest Request Object
type ExecuteLoginConnectionNewRequest struct {

	// 连接ID
	ConnectionId string `json:"connection_id"`

	Body *ExecuteLoginConnectionNewRequestBody `json:"body,omitempty"`
}

func (o ExecuteLoginConnectionNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteLoginConnectionNewRequest struct{}"
	}

	return strings.Join([]string{"ExecuteLoginConnectionNewRequest", string(data)}, " ")
}
