package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateConnectionResponse Response Object
type CreateConnectionResponse struct {

	// 连接ID
	ConnectionId   *string `json:"connection_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConnectionResponse struct{}"
	}

	return strings.Join([]string{"CreateConnectionResponse", string(data)}, " ")
}
