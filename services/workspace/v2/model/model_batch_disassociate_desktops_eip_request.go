package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDisassociateDesktopsEipRequest Request Object
type BatchDisassociateDesktopsEipRequest struct {
	Body *BatchDisassociateDesktopsEipReq `json:"body,omitempty"`
}

func (o BatchDisassociateDesktopsEipRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDisassociateDesktopsEipRequest struct{}"
	}

	return strings.Join([]string{"BatchDisassociateDesktopsEipRequest", string(data)}, " ")
}
