package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteProtocolConfigResponse Response Object
type DeleteProtocolConfigResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteProtocolConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteProtocolConfigResponse struct{}"
	}

	return strings.Join([]string{"DeleteProtocolConfigResponse", string(data)}, " ")
}
