package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateChannelsResponse Response Object
type BatchCreateChannelsResponse struct {

	// 操作记录id
	OperationId    *string `json:"operation_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchCreateChannelsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateChannelsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateChannelsResponse", string(data)}, " ")
}
