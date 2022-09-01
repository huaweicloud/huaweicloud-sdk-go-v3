package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DisassociatePublicipsRequest struct {

	// 弹性公网IP的ID
	PublicipId string `json:"publicip_id" xml:"publicip_id"`

	Body *DisassociatePublicipsRequestBody `json:"body,omitempty" xml:"body"`
}

func (o DisassociatePublicipsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociatePublicipsRequest struct{}"
	}

	return strings.Join([]string{"DisassociatePublicipsRequest", string(data)}, " ")
}
