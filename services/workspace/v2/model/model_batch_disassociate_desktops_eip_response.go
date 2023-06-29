package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDisassociateDesktopsEipResponse Response Object
type BatchDisassociateDesktopsEipResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDisassociateDesktopsEipResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDisassociateDesktopsEipResponse struct{}"
	}

	return strings.Join([]string{"BatchDisassociateDesktopsEipResponse", string(data)}, " ")
}
