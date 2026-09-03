package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLatestDeadLockSnapshot4ApiRequest Request Object
type ShowLatestDeadLockSnapshot4ApiRequest struct {

	// 连接ID
	ConnectionId string `json:"connection_id"`

	// 死锁id
	Id *int32 `json:"id,omitempty"`
}

func (o ShowLatestDeadLockSnapshot4ApiRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLatestDeadLockSnapshot4ApiRequest struct{}"
	}

	return strings.Join([]string{"ShowLatestDeadLockSnapshot4ApiRequest", string(data)}, " ")
}
