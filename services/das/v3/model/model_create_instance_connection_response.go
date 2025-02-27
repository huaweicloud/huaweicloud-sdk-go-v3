package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInstanceConnectionResponse Response Object
type CreateInstanceConnectionResponse struct {

	// 连接编号。
	ConnectionId   *string `json:"connection_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateInstanceConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceConnectionResponse struct{}"
	}

	return strings.Join([]string{"CreateInstanceConnectionResponse", string(data)}, " ")
}
