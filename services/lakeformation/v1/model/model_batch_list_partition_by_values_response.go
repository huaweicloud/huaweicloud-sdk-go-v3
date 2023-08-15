package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchListPartitionByValuesResponse Response Object
type BatchListPartitionByValuesResponse struct {
	Body           *[]Partition `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o BatchListPartitionByValuesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchListPartitionByValuesResponse struct{}"
	}

	return strings.Join([]string{"BatchListPartitionByValuesResponse", string(data)}, " ")
}
