package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisassociateRouterRequest Request Object
type DisassociateRouterRequest struct {

	// 域名ID。
	ZoneId string `json:"zone_id"`

	Body *DisassociaterouterRequestBody `json:"body,omitempty"`
}

func (o DisassociateRouterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateRouterRequest struct{}"
	}

	return strings.Join([]string{"DisassociateRouterRequest", string(data)}, " ")
}
