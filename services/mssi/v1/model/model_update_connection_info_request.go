package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateConnectionInfoRequest Request Object
type UpdateConnectionInfoRequest struct {

	// ID of Connection
	ConnectId string `json:"connect_id"`

	Body *ConnectionInfo `json:"body,omitempty"`
}

func (o UpdateConnectionInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConnectionInfoRequest struct{}"
	}

	return strings.Join([]string{"UpdateConnectionInfoRequest", string(data)}, " ")
}
