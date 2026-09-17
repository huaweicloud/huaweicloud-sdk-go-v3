package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InvokeKubeApiResponse Response Object
type InvokeKubeApiResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o InvokeKubeApiResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InvokeKubeApiResponse struct{}"
	}

	return strings.Join([]string{"InvokeKubeApiResponse", string(data)}, " ")
}
