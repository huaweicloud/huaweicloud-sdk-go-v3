package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPolicyInstancesRequest Request Object
type ListPolicyInstancesRequest struct {
}

func (o ListPolicyInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPolicyInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListPolicyInstancesRequest", string(data)}, " ")
}
