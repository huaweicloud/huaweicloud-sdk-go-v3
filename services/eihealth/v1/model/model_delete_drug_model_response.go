package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDrugModelResponse Response Object
type DeleteDrugModelResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteDrugModelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDrugModelResponse struct{}"
	}

	return strings.Join([]string{"DeleteDrugModelResponse", string(data)}, " ")
}
