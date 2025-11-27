package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListReplicationErrorsResponse Response Object
type ListReplicationErrorsResponse struct {

	// 报错列表。
	Errors *[]ListReplicationErrorsResult `json:"errors,omitempty"`

	// 报错总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListReplicationErrorsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListReplicationErrorsResponse struct{}"
	}

	return strings.Join([]string{"ListReplicationErrorsResponse", string(data)}, " ")
}
