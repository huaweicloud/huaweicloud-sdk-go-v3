package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDrugModelResponse Response Object
type UpdateDrugModelResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateDrugModelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDrugModelResponse struct{}"
	}

	return strings.Join([]string{"UpdateDrugModelResponse", string(data)}, " ")
}
