package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDedicatedHostTypesResponse Response Object
type ListDedicatedHostTypesResponse struct {

	// 可用的专属主机类型。
	DedicatedHostTypes *[]RespHostType `json:"dedicated_host_types,omitempty"`
	HttpStatusCode     int             `json:"-"`
}

func (o ListDedicatedHostTypesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDedicatedHostTypesResponse struct{}"
	}

	return strings.Join([]string{"ListDedicatedHostTypesResponse", string(data)}, " ")
}
