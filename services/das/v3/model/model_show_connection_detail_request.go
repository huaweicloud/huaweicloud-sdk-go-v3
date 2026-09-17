package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowConnectionDetailRequest Request Object
type ShowConnectionDetailRequest struct {

	// 连接ID
	ConnectionId string `json:"connection_id"`
}

func (o ShowConnectionDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConnectionDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowConnectionDetailRequest", string(data)}, " ")
}
