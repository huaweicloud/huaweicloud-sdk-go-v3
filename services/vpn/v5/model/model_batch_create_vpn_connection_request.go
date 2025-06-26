package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateVpnConnectionRequest Request Object
type BatchCreateVpnConnectionRequest struct {
	Body *BatchCreateVpnConnectionRequestBody `json:"body,omitempty"`
}

func (o BatchCreateVpnConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateVpnConnectionRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateVpnConnectionRequest", string(data)}, " ")
}
