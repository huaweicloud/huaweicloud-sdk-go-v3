package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDrugModelRequest Request Object
type DeleteDrugModelRequest struct {

	// 模型id
	ModelId string `json:"model_id"`
}

func (o DeleteDrugModelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDrugModelRequest struct{}"
	}

	return strings.Join([]string{"DeleteDrugModelRequest", string(data)}, " ")
}
