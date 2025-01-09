package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLoginStateResponse Response Object
type ListLoginStateResponse struct {

	// 使用中。
	InUseNum *int32 `json:"in_use_num,omitempty"`

	// 关机数目(关机中、已关机)。
	StopNum *int32 `json:"stop_num,omitempty"`

	// 未注册数目。
	UnregisteredNum *int32 `json:"unregistered_num,omitempty"`

	// 未注册数目。
	UnableToConnectNum *int32 `json:"unable_to_connect_num,omitempty"`

	// 空闲数目。
	ReadyNum *int32 `json:"ready_num,omitempty"`

	// 断开连接数目。
	DisconnectedNum *int32 `json:"disconnected_num,omitempty"`
	HttpStatusCode  int    `json:"-"`
}

func (o ListLoginStateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLoginStateResponse struct{}"
	}

	return strings.Join([]string{"ListLoginStateResponse", string(data)}, " ")
}
