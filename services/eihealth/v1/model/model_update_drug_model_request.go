package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDrugModelRequest Request Object
type UpdateDrugModelRequest struct {

	// 模型id
	ModelId string `json:"model_id"`

	Body *UpdateDrugModelReq `json:"body,omitempty"`
}

func (o UpdateDrugModelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDrugModelRequest struct{}"
	}

	return strings.Join([]string{"UpdateDrugModelRequest", string(data)}, " ")
}
