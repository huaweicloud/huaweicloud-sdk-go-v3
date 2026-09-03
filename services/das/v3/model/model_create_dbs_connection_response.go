package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDbsConnectionResponse Response Object
type CreateDbsConnectionResponse struct {

	// 连接ID
	ConnectionId   *string `json:"connection_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDbsConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDbsConnectionResponse struct{}"
	}

	return strings.Join([]string{"CreateDbsConnectionResponse", string(data)}, " ")
}
