package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTagPolicyServicesResponse Response Object
type ListTagPolicyServicesResponse struct {
	Services       *[]TagPolicyServiceDto `json:"services,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListTagPolicyServicesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTagPolicyServicesResponse struct{}"
	}

	return strings.Join([]string{"ListTagPolicyServicesResponse", string(data)}, " ")
}
