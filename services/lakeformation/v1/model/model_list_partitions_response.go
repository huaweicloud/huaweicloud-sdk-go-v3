package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListPartitionsResponse struct {
	Partitions *[]Partition `json:"partitions,omitempty"`

	PageInfo       *PagedInfo `json:"page_info,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListPartitionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPartitionsResponse struct{}"
	}

	return strings.Join([]string{"ListPartitionsResponse", string(data)}, " ")
}
