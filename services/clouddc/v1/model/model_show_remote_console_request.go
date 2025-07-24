package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRemoteConsoleRequest Request Object
type ShowRemoteConsoleRequest struct {

	// imetal server id
	Id string `json:"id"`
}

func (o ShowRemoteConsoleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRemoteConsoleRequest struct{}"
	}

	return strings.Join([]string{"ShowRemoteConsoleRequest", string(data)}, " ")
}
