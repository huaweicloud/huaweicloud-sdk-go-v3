package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteVirsubnetCidrReservationRequest Request Object
type DeleteVirsubnetCidrReservationRequest struct {

	// **参数解释**： 子网预留网段的资源ID。 **取值范围**： 不涉及。
	VirsubnetCidrReservationId string `json:"virsubnet_cidr_reservation_id"`
}

func (o DeleteVirsubnetCidrReservationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVirsubnetCidrReservationRequest struct{}"
	}

	return strings.Join([]string{"DeleteVirsubnetCidrReservationRequest", string(data)}, " ")
}
