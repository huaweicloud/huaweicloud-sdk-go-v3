package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVirsubnetCidrReservationRequest Request Object
type CreateVirsubnetCidrReservationRequest struct {
	Body *CreateVirsubnetCidrReservationRequestBody `json:"body,omitempty"`
}

func (o CreateVirsubnetCidrReservationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVirsubnetCidrReservationRequest struct{}"
	}

	return strings.Join([]string{"CreateVirsubnetCidrReservationRequest", string(data)}, " ")
}
