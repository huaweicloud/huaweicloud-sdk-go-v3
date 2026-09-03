package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteSplitSqlRequest Request Object
type ExecuteSplitSqlRequest struct {

	// 连接ID
	ConnectionId string `json:"connection_id"`

	Body *ExecuteSplitSqlRequestBody `json:"body,omitempty"`
}

func (o ExecuteSplitSqlRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteSplitSqlRequest struct{}"
	}

	return strings.Join([]string{"ExecuteSplitSqlRequest", string(data)}, " ")
}
