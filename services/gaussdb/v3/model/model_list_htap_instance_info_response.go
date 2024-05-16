package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHtapInstanceInfoResponse Response Object
type ListHtapInstanceInfoResponse struct {

	// HTAP实例个数。
	Total *int32 `json:"total,omitempty"`

	// HTAP实例信息。
	Instances *[]HtapInstanceListInstances `json:"instances,omitempty"`

	// 最大HTAP实例个数。
	MaxHtapInstanceNumOfTaurus *int32 `json:"max_htap_instance_num_of_taurus,omitempty"`
	HttpStatusCode             int    `json:"-"`
}

func (o ListHtapInstanceInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHtapInstanceInfoResponse struct{}"
	}

	return strings.Join([]string{"ListHtapInstanceInfoResponse", string(data)}, " ")
}
