package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceDetailsResponse Response Object
type ListInstanceDetailsResponse struct {

	// 实例信息。
	Instances *[]ListInstancesResult `json:"instances,omitempty"`

	// 总记录数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListInstanceDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceDetailsResponse", string(data)}, " ")
}
