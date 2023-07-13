package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchRemoveOrgsFromChannelResponse Response Object
type BatchRemoveOrgsFromChannelResponse struct {

	// 操作记录id
	OperationId    *string `json:"operation_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchRemoveOrgsFromChannelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRemoveOrgsFromChannelResponse struct{}"
	}

	return strings.Join([]string{"BatchRemoveOrgsFromChannelResponse", string(data)}, " ")
}
