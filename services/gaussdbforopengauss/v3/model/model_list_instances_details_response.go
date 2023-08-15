package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstancesDetailsResponse Response Object
type ListInstancesDetailsResponse struct {

	// 实例信息。
	Instances *[]ListInstanceResult `json:"instances,omitempty"`

	// 总记录数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListInstancesDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListInstancesDetailsResponse", string(data)}, " ")
}
