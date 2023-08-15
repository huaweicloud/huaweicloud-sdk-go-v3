package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDliAgencyRequest Request Object
type CreateDliAgencyRequest struct {
	Body *CreateAgencyRequest `json:"body,omitempty"`
}

func (o CreateDliAgencyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDliAgencyRequest struct{}"
	}

	return strings.Join([]string{"CreateDliAgencyRequest", string(data)}, " ")
}
