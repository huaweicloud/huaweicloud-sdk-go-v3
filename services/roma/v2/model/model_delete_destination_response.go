package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDestinationResponse Response Object
type DeleteDestinationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteDestinationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDestinationResponse struct{}"
	}

	return strings.Join([]string{"DeleteDestinationResponse", string(data)}, " ")
}
