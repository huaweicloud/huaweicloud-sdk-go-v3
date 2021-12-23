package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteCheckpointResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteCheckpointResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCheckpointResponse struct{}"
	}

	return strings.Join([]string{"DeleteCheckpointResponse", string(data)}, " ")
}
