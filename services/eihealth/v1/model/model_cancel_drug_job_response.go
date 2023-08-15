package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelDrugJobResponse Response Object
type CancelDrugJobResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CancelDrugJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelDrugJobResponse struct{}"
	}

	return strings.Join([]string{"CancelDrugJobResponse", string(data)}, " ")
}
