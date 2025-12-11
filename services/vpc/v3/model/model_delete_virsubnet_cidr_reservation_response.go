package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteVirsubnetCidrReservationResponse Response Object
type DeleteVirsubnetCidrReservationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteVirsubnetCidrReservationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVirsubnetCidrReservationResponse struct{}"
	}

	return strings.Join([]string{"DeleteVirsubnetCidrReservationResponse", string(data)}, " ")
}
