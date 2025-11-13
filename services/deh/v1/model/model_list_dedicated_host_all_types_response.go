package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDedicatedHostAllTypesResponse Response Object
type ListDedicatedHostAllTypesResponse struct {

	// 指定可用的DeH类型。
	DedicatedHostTypes *[]DedicatedHostType `json:"dedicated_host_types,omitempty"`
	HttpStatusCode     int                  `json:"-"`
}

func (o ListDedicatedHostAllTypesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDedicatedHostAllTypesResponse struct{}"
	}

	return strings.Join([]string{"ListDedicatedHostAllTypesResponse", string(data)}, " ")
}
