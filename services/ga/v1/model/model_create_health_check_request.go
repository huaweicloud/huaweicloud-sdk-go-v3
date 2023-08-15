package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateHealthCheckRequest Request Object
type CreateHealthCheckRequest struct {
	Body *CreateHealthCheckRequestBody `json:"body,omitempty"`
}

func (o CreateHealthCheckRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHealthCheckRequest struct{}"
	}

	return strings.Join([]string{"CreateHealthCheckRequest", string(data)}, " ")
}
