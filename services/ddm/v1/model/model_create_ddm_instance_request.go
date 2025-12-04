package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDdmInstanceRequest Request Object
type CreateDdmInstanceRequest struct {
	Body *CreateDdmInstanceReq `json:"body,omitempty"`
}

func (o CreateDdmInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDdmInstanceRequest struct{}"
	}

	return strings.Join([]string{"CreateDdmInstanceRequest", string(data)}, " ")
}
