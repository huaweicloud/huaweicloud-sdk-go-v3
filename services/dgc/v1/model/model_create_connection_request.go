package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateConnectionRequest struct {
	Body *ConnectionInfo `json:"body,omitempty"`
}

func (o CreateConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConnectionRequest struct{}"
	}

	return strings.Join([]string{"CreateConnectionRequest", string(data)}, " ")
}
