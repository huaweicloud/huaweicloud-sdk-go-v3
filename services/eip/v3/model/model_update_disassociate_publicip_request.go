package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDisassociatePublicipRequest Request Object
type UpdateDisassociatePublicipRequest struct {

	// 弹性公网IP的ID
	PublicipId string `json:"publicip_id"`

	Body *DisassociatePublicipsRequestBody `json:"body,omitempty"`
}

func (o UpdateDisassociatePublicipRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDisassociatePublicipRequest struct{}"
	}

	return strings.Join([]string{"UpdateDisassociatePublicipRequest", string(data)}, " ")
}
