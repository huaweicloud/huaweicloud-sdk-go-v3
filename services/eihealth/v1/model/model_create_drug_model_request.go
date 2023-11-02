package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDrugModelRequest Request Object
type CreateDrugModelRequest struct {
	Body *CreateModelReq `json:"body,omitempty"`
}

func (o CreateDrugModelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDrugModelRequest struct{}"
	}

	return strings.Join([]string{"CreateDrugModelRequest", string(data)}, " ")
}
