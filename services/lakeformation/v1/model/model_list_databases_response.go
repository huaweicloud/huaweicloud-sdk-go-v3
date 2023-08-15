package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDatabasesResponse Response Object
type ListDatabasesResponse struct {
	Databases *[]Database `json:"databases,omitempty"`

	PageInfo       *PagedInfo `json:"page_info,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListDatabasesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatabasesResponse struct{}"
	}

	return strings.Join([]string{"ListDatabasesResponse", string(data)}, " ")
}
