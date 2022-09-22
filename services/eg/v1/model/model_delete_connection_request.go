package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteConnectionRequest struct {

	// 指定查询的目标连接ID
	ConnectionId string `json:"connection_id"`
}

func (o DeleteConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteConnectionRequest struct{}"
	}

	return strings.Join([]string{"DeleteConnectionRequest", string(data)}, " ")
}
