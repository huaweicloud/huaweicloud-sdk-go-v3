package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDetailOfConnectionRequest Request Object
type ShowDetailOfConnectionRequest struct {

	// 指定查询的目标连接ID
	ConnectionId string `json:"connection_id"`
}

func (o ShowDetailOfConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDetailOfConnectionRequest struct{}"
	}

	return strings.Join([]string{"ShowDetailOfConnectionRequest", string(data)}, " ")
}
