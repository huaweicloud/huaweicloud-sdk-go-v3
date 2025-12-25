package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCollectorConnectionResponse Response Object
type UpdateCollectorConnectionResponse struct {

	// UUID
	ConnectionId   *string `json:"connection_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateCollectorConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCollectorConnectionResponse struct{}"
	}

	return strings.Join([]string{"UpdateCollectorConnectionResponse", string(data)}, " ")
}
