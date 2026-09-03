package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelConnectionProcessRequest Request Object
type CancelConnectionProcessRequest struct {

	// 连接ID
	ConnectionId string `json:"connection_id"`

	Body *CancelConnectionProcessRequestBody `json:"body,omitempty"`
}

func (o CancelConnectionProcessRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelConnectionProcessRequest struct{}"
	}

	return strings.Join([]string{"CancelConnectionProcessRequest", string(data)}, " ")
}
