package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVirtualInterfaceRequest Request Object
type CreateVirtualInterfaceRequest struct {
	Body *CreateVirtualInterfaceRequestBody `json:"body,omitempty"`
}

func (o CreateVirtualInterfaceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVirtualInterfaceRequest struct{}"
	}

	return strings.Join([]string{"CreateVirtualInterfaceRequest", string(data)}, " ")
}
