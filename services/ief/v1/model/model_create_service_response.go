package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateServiceResponse struct {
	Service        *ServiceRespDetail `json:"service,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o CreateServiceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateServiceResponse struct{}"
	}

	return strings.Join([]string{"CreateServiceResponse", string(data)}, " ")
}
