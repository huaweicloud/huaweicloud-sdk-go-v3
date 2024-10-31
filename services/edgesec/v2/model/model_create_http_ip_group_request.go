package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateHttpIpGroupRequest Request Object
type CreateHttpIpGroupRequest struct {
	Body *CreateHttpIpGroupRequestBody `json:"body,omitempty"`
}

func (o CreateHttpIpGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHttpIpGroupRequest struct{}"
	}

	return strings.Join([]string{"CreateHttpIpGroupRequest", string(data)}, " ")
}
