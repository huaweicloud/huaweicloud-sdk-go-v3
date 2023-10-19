package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPartitionNamesWithoutLimitResponse Response Object
type ListPartitionNamesWithoutLimitResponse struct {
	Body           *[]string `json:"body,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListPartitionNamesWithoutLimitResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPartitionNamesWithoutLimitResponse struct{}"
	}

	return strings.Join([]string{"ListPartitionNamesWithoutLimitResponse", string(data)}, " ")
}
