package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SyncConnectionsNewResponse Response Object
type SyncConnectionsNewResponse struct {

	// 同步信息
	Msg            *string `json:"msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SyncConnectionsNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SyncConnectionsNewResponse struct{}"
	}

	return strings.Join([]string{"SyncConnectionsNewResponse", string(data)}, " ")
}
