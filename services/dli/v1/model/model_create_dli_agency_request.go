package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDliAgencyRequest Request Object
type CreateDliAgencyRequest struct {
	Body *CreateDliAgencyRequestBody `json:"body,omitempty"`
}

func (o CreateDliAgencyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDliAgencyRequest struct{}"
	}

	return strings.Join([]string{"CreateDliAgencyRequest", string(data)}, " ")
}
