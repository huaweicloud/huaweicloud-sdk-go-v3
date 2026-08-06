package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConnectionIdsItem struct {

	// 连接ID
	ConnectionId *string `json:"connection_id,omitempty"`
}

func (o ConnectionIdsItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConnectionIdsItem struct{}"
	}

	return strings.Join([]string{"ConnectionIdsItem", string(data)}, " ")
}
