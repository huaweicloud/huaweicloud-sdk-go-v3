package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteBasicPluginResponse Response Object
type DeleteBasicPluginResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteBasicPluginResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBasicPluginResponse struct{}"
	}

	return strings.Join([]string{"DeleteBasicPluginResponse", string(data)}, " ")
}
