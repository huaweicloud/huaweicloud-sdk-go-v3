package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVirsubnetCidrReservationResponse Response Object
type ShowVirsubnetCidrReservationResponse struct {

	// **参数解释**： 请求ID。 **取值范围**： 不涉及。
	RequestId *string `json:"request_id,omitempty"`

	VirsubnetCidrReservation *VirsubnetCidrReservation `json:"virsubnet_cidr_reservation,omitempty"`
	HttpStatusCode           int                       `json:"-"`
}

func (o ShowVirsubnetCidrReservationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVirsubnetCidrReservationResponse struct{}"
	}

	return strings.Join([]string{"ShowVirsubnetCidrReservationResponse", string(data)}, " ")
}
