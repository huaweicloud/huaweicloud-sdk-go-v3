package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetachPluginFromApiResponse Response Object
type DetachPluginFromApiResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DetachPluginFromApiResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetachPluginFromApiResponse struct{}"
	}

	return strings.Join([]string{"DetachPluginFromApiResponse", string(data)}, " ")
}
