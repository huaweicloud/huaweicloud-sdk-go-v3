package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DisassociateRouterRequest struct {
	// 待解关联zone的ID。

	ZoneId string `json:"zone_id"`

	Body *DisassociaterouterReq `json:"body,omitempty"`
}

func (o DisassociateRouterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateRouterRequest struct{}"
	}

	return strings.Join([]string{"DisassociateRouterRequest", string(data)}, " ")
}
