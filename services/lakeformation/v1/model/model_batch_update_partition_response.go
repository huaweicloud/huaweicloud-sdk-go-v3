package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdatePartitionResponse Response Object
type BatchUpdatePartitionResponse struct {
	Body           *[]Partition `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o BatchUpdatePartitionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdatePartitionResponse struct{}"
	}

	return strings.Join([]string{"BatchUpdatePartitionResponse", string(data)}, " ")
}
