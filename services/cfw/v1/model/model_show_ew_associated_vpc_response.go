package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEwAssociatedVpcResponse Response Object
type ShowEwAssociatedVpcResponse struct {
	Data           *ShowEwAssociatedVpcRespData `json:"data,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ShowEwAssociatedVpcResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEwAssociatedVpcResponse struct{}"
	}

	return strings.Join([]string{"ShowEwAssociatedVpcResponse", string(data)}, " ")
}
