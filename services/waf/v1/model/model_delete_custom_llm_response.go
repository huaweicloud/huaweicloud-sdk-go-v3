package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCustomLlmResponse Response Object
type DeleteCustomLlmResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteCustomLlmResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCustomLlmResponse struct{}"
	}

	return strings.Join([]string{"DeleteCustomLlmResponse", string(data)}, " ")
}
