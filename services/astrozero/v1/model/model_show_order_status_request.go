package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowOrderStatusRequest Request Object
type ShowOrderStatusRequest struct {
	Body *ShowOrderStatusReq `json:"body,omitempty"`
}

func (o ShowOrderStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOrderStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowOrderStatusRequest", string(data)}, " ")
}
