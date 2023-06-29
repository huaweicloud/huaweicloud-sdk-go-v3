package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnfreezeNodeResponse Response Object
type UnfreezeNodeResponse struct {

	// 设备ID
	NodeId         *string `json:"node_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UnfreezeNodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnfreezeNodeResponse struct{}"
	}

	return strings.Join([]string{"UnfreezeNodeResponse", string(data)}, " ")
}
