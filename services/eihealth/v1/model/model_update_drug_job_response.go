package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDrugJobResponse Response Object
type UpdateDrugJobResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateDrugJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDrugJobResponse struct{}"
	}

	return strings.Join([]string{"UpdateDrugJobResponse", string(data)}, " ")
}
