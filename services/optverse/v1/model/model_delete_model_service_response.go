package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteModelServiceResponse Response Object
type DeleteModelServiceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteModelServiceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteModelServiceResponse struct{}"
	}

	return strings.Join([]string{"DeleteModelServiceResponse", string(data)}, " ")
}
