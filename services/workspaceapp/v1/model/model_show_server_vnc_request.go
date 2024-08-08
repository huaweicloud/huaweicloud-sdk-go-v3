package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowServerVncRequest Request Object
type ShowServerVncRequest struct {

	// 服务器唯一标识。
	ServerId string `json:"server_id"`
}

func (o ShowServerVncRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServerVncRequest struct{}"
	}

	return strings.Join([]string{"ShowServerVncRequest", string(data)}, " ")
}
