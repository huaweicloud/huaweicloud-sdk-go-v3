package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVgwRequest Request Object
type ShowVgwRequest struct {

	// VPN网关实例ID
	VgwId string `json:"vgw_id"`
}

func (o ShowVgwRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVgwRequest struct{}"
	}

	return strings.Join([]string{"ShowVgwRequest", string(data)}, " ")
}
