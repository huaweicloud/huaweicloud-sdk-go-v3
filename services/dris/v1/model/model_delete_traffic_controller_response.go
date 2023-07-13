package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTrafficControllerResponse Response Object
type DeleteTrafficControllerResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteTrafficControllerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTrafficControllerResponse struct{}"
	}

	return strings.Join([]string{"DeleteTrafficControllerResponse", string(data)}, " ")
}
