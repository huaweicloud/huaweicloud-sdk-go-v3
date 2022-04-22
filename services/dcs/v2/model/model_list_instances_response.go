package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListInstancesResponse struct {

	// 实例个数。
	InstanceNum *int32 `json:"instance_num,omitempty"`

	// 实例的详情数组。
	Instances      *[]InstanceListInfo `json:"instances,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListInstancesResponse", string(data)}, " ")
}
