package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteServiceInstanceResponse Response Object
type DeleteServiceInstanceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteServiceInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteServiceInstanceResponse struct{}"
	}

	return strings.Join([]string{"DeleteServiceInstanceResponse", string(data)}, " ")
}
