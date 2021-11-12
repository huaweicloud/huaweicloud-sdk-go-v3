package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteWindowsBareMetalServerPasswordResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteWindowsBareMetalServerPasswordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteWindowsBareMetalServerPasswordResponse struct{}"
	}

	return strings.Join([]string{"DeleteWindowsBareMetalServerPasswordResponse", string(data)}, " ")
}
