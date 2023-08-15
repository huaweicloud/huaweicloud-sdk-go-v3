package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTagPolicyServicesRequest Request Object
type ListTagPolicyServicesRequest struct {
}

func (o ListTagPolicyServicesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTagPolicyServicesRequest struct{}"
	}

	return strings.Join([]string{"ListTagPolicyServicesRequest", string(data)}, " ")
}
