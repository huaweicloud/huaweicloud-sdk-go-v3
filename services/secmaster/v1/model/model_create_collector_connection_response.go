package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCollectorConnectionResponse Response Object
type CreateCollectorConnectionResponse struct {

	// UUID
	ConnectionId   *string `json:"connection_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateCollectorConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCollectorConnectionResponse struct{}"
	}

	return strings.Join([]string{"CreateCollectorConnectionResponse", string(data)}, " ")
}
