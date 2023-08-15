package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteWindowsBareMetalServerPasswordResponse Response Object
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
