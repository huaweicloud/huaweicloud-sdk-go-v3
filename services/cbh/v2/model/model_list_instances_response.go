package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstancesResponse Response Object
type ListInstancesResponse struct {

	// 云堡垒机实例总数。
	Total *int32 `json:"total,omitempty"`

	// 云堡垒机实例列表信息。
	Instance       *[]InstanceDetail `json:"instance,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListInstancesResponse", string(data)}, " ")
}
