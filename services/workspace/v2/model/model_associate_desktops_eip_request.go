package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateDesktopsEipRequest Request Object
type AssociateDesktopsEipRequest struct {
	Body *AssociateDesktopsEipReq `json:"body,omitempty"`
}

func (o AssociateDesktopsEipRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateDesktopsEipRequest struct{}"
	}

	return strings.Join([]string{"AssociateDesktopsEipRequest", string(data)}, " ")
}
