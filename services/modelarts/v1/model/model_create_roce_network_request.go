package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRoceNetworkRequest Request Object
type CreateRoceNetworkRequest struct {
	Body *ServerRoceNetworkRequest `json:"body,omitempty"`
}

func (o CreateRoceNetworkRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRoceNetworkRequest struct{}"
	}

	return strings.Join([]string{"CreateRoceNetworkRequest", string(data)}, " ")
}
