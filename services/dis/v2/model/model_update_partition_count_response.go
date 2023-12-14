package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePartitionCountResponse Response Object
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
