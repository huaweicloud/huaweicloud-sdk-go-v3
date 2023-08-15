package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteVehicleResponse Response Object
type DeleteVehicleResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteVehicleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVehicleResponse struct{}"
	}

	return strings.Join([]string{"DeleteVehicleResponse", string(data)}, " ")
}
