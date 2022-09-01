package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type NovaDisassociateSecurityGroupRequest struct {

	// 云服务器ID。
	ServerId string `json:"server_id" xml:"server_id"`

	Body *NovaDisassociateSecurityGroupRequestBody `json:"body,omitempty" xml:"body"`
}

func (o NovaDisassociateSecurityGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaDisassociateSecurityGroupRequest struct{}"
	}

	return strings.Join([]string{"NovaDisassociateSecurityGroupRequest", string(data)}, " ")
}
