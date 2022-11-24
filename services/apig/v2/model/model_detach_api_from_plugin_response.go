package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DetachApiFromPluginResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DetachApiFromPluginResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetachApiFromPluginResponse struct{}"
	}

	return strings.Join([]string{"DetachApiFromPluginResponse", string(data)}, " ")
}
