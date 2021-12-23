package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdatePartitionCountResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdatePartitionCountResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePartitionCountResponse struct{}"
	}

	return strings.Join([]string{"UpdatePartitionCountResponse", string(data)}, " ")
}
