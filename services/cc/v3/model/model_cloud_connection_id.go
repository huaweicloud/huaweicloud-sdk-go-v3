package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CloudConnectionId 云连接实例ID。
type CloudConnectionId struct {

	// 云连接实例ID。
	CloudConnectionId string `json:"cloud_connection_id"`
}

func (o CloudConnectionId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloudConnectionId struct{}"
	}

	return strings.Join([]string{"CloudConnectionId", string(data)}, " ")
}
