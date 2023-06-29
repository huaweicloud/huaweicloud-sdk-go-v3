package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTablesResponse Response Object
type ListTablesResponse struct {
	Tables *[]Table `json:"tables,omitempty"`

	PageInfo       *PagedInfo `json:"page_info,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListTablesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTablesResponse struct{}"
	}

	return strings.Join([]string{"ListTablesResponse", string(data)}, " ")
}
