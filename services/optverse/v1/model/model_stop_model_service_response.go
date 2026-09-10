package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopModelServiceResponse Response Object
type StopModelServiceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StopModelServiceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopModelServiceResponse struct{}"
	}

	return strings.Join([]string{"StopModelServiceResponse", string(data)}, " ")
}
