package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateServiceRequest Request Object
type CreateServiceRequest struct {
	Body *ServiceRequestBody `json:"body,omitempty"`
}

func (o CreateServiceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateServiceRequest struct{}"
	}

	return strings.Join([]string{"CreateServiceRequest", string(data)}, " ")
}
