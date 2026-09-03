package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteFormatSqlRequest Request Object
type ExecuteFormatSqlRequest struct {

	// 连接ID
	ConnectionId string `json:"connection_id"`

	Body *ExecuteFormatSqlRequestBody `json:"body,omitempty"`
}

func (o ExecuteFormatSqlRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteFormatSqlRequest struct{}"
	}

	return strings.Join([]string{"ExecuteFormatSqlRequest", string(data)}, " ")
}
