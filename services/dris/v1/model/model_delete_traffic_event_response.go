package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteTrafficEventResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteTrafficEventResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTrafficEventResponse struct{}"
	}

	return strings.Join([]string{"DeleteTrafficEventResponse", string(data)}, " ")
}
