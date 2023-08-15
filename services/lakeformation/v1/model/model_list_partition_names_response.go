package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPartitionNamesResponse Response Object
type ListPartitionNamesResponse struct {

	// 分区名字列表
	PartitionNameList *[]string `json:"partition_name_list,omitempty"`

	PageInfo       *PagedInfo `json:"page_info,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListPartitionNamesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPartitionNamesResponse struct{}"
	}

	return strings.Join([]string{"ListPartitionNamesResponse", string(data)}, " ")
}
