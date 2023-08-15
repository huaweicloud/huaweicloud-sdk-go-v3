package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCommandRequest Request Object
type ShowCommandRequest struct {

	// 命令对应的服务器ID
	ServerId string `json:"server_id"`
}

func (o ShowCommandRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCommandRequest struct{}"
	}

	return strings.Join([]string{"ShowCommandRequest", string(data)}, " ")
}
