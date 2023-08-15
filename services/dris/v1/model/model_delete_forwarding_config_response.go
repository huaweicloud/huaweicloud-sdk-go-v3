package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteForwardingConfigResponse Response Object
type DeleteForwardingConfigResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteForwardingConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteForwardingConfigResponse struct{}"
	}

	return strings.Join([]string{"DeleteForwardingConfigResponse", string(data)}, " ")
}
