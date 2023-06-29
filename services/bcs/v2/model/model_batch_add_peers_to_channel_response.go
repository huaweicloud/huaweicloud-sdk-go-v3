package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchAddPeersToChannelResponse Response Object
type BatchAddPeersToChannelResponse struct {

	// 操作记录id
	OperationId    *string `json:"operation_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchAddPeersToChannelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddPeersToChannelResponse struct{}"
	}

	return strings.Join([]string{"BatchAddPeersToChannelResponse", string(data)}, " ")
}
