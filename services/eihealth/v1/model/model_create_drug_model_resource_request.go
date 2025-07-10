package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDrugModelResourceRequest Request Object
type CreateDrugModelResourceRequest struct {
	Body *CreateDrugModelResourceReq `json:"body,omitempty"`
}

func (o CreateDrugModelResourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDrugModelResourceRequest struct{}"
	}

	return strings.Join([]string{"CreateDrugModelResourceRequest", string(data)}, " ")
}
