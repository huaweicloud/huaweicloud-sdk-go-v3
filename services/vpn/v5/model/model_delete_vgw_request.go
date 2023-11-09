package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteVgwRequest Request Object
type DeleteVgwRequest struct {

	// VPN网关实例ID
	VgwId string `json:"vgw_id"`
}

func (o DeleteVgwRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVgwRequest struct{}"
	}

	return strings.Join([]string{"DeleteVgwRequest", string(data)}, " ")
}
