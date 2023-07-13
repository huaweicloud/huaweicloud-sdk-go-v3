package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeletePartitionResponse Response Object
type BatchDeletePartitionResponse struct {
	Body           *[]Partition `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o BatchDeletePartitionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeletePartitionResponse struct{}"
	}

	return strings.Join([]string{"BatchDeletePartitionResponse", string(data)}, " ")
}
