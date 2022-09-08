package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreatePublicIpRequest struct {
	Body *CreatePublicIpRequestBody `json:"body,omitempty"`
}

func (o CreatePublicIpRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePublicIpRequest struct{}"
	}

	return strings.Join([]string{"CreatePublicIpRequest", string(data)}, " ")
}
