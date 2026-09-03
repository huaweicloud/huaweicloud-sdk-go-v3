package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowExecuteResultWithoutKeyRequest Request Object
type ShowExecuteResultWithoutKeyRequest struct {

	// 连接ID
	ConnectionId string `json:"connection_id"`

	// 执行ID
	ExecuteId *string `json:"execute_id,omitempty"`
}

func (o ShowExecuteResultWithoutKeyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowExecuteResultWithoutKeyRequest struct{}"
	}

	return strings.Join([]string{"ShowExecuteResultWithoutKeyRequest", string(data)}, " ")
}
