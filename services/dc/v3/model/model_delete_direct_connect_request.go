package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDirectConnectRequest Request Object
type DeleteDirectConnectRequest struct {

	// 物理专线连接ID。
	DirectConnectId string `json:"direct_connect_id"`
}

func (o DeleteDirectConnectRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDirectConnectRequest struct{}"
	}

	return strings.Join([]string{"DeleteDirectConnectRequest", string(data)}, " ")
}
