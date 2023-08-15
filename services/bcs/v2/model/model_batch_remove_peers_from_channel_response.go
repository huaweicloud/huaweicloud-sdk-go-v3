package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchRemovePeersFromChannelResponse Response Object
type BatchRemovePeersFromChannelResponse struct {

	// 操作记录id
	OperationId    *string `json:"operation_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchRemovePeersFromChannelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRemovePeersFromChannelResponse struct{}"
	}

	return strings.Join([]string{"BatchRemovePeersFromChannelResponse", string(data)}, " ")
}
