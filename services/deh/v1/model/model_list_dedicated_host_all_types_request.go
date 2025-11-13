package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDedicatedHostAllTypesRequest Request Object
type ListDedicatedHostAllTypesRequest struct {
}

func (o ListDedicatedHostAllTypesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDedicatedHostAllTypesRequest struct{}"
	}

	return strings.Join([]string{"ListDedicatedHostAllTypesRequest", string(data)}, " ")
}
