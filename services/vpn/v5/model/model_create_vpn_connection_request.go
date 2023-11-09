package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVpnConnectionRequest Request Object
type CreateVpnConnectionRequest struct {
	Body *CreateVpnConnectionRequestBody `json:"body,omitempty"`
}

func (o CreateVpnConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVpnConnectionRequest struct{}"
	}

	return strings.Join([]string{"CreateVpnConnectionRequest", string(data)}, " ")
}
