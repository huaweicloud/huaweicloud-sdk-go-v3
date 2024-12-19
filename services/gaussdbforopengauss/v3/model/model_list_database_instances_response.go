package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDatabaseInstancesResponse Response Object
type ListDatabaseInstancesResponse struct {

	// 实例信息。
	Instances *[]ListInstancesResult `json:"instances,omitempty"`

	// 总记录数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDatabaseInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatabaseInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListDatabaseInstancesResponse", string(data)}, " ")
}
