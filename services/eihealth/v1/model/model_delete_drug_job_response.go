package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteDrugJobResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteDrugJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDrugJobResponse struct{}"
	}

	return strings.Join([]string{"DeleteDrugJobResponse", string(data)}, " ")
}
