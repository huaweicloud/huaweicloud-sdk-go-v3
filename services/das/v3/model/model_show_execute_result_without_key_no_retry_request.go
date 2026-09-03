package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowExecuteResultWithoutKeyNoRetryRequest Request Object
type ShowExecuteResultWithoutKeyNoRetryRequest struct {

	// 连接ID
	ConnectionId string `json:"connection_id"`

	Body *ShowExecuteResultWithoutKeyNoRetryRequestBody `json:"body,omitempty"`
}

func (o ShowExecuteResultWithoutKeyNoRetryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowExecuteResultWithoutKeyNoRetryRequest struct{}"
	}

	return strings.Join([]string{"ShowExecuteResultWithoutKeyNoRetryRequest", string(data)}, " ")
}
