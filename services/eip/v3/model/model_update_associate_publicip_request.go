package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAssociatePublicipRequest Request Object
type UpdateAssociatePublicipRequest struct {

	// 弹性公网IP的ID
	PublicipId string `json:"publicip_id"`

	Body *AssociatePublicipsRequestBody `json:"body,omitempty"`
}

func (o UpdateAssociatePublicipRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAssociatePublicipRequest struct{}"
	}

	return strings.Join([]string{"UpdateAssociatePublicipRequest", string(data)}, " ")
}
