package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateVirsubnetCidrReservationRequest Request Object
type UpdateVirsubnetCidrReservationRequest struct {

	// **参数解释**： 子网预留网段的资源ID。 **取值范围**： 不涉及。
	VirsubnetCidrReservationId string `json:"virsubnet_cidr_reservation_id"`

	Body *UpdateVirsubnetCidrReservationRequestBody `json:"body,omitempty"`
}

func (o UpdateVirsubnetCidrReservationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVirsubnetCidrReservationRequest struct{}"
	}

	return strings.Join([]string{"UpdateVirsubnetCidrReservationRequest", string(data)}, " ")
}
