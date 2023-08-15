package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FreezeNodeResponse Response Object
type FreezeNodeResponse struct {

	// 设备ID
	NodeId         *string `json:"node_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o FreezeNodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FreezeNodeResponse struct{}"
	}

	return strings.Join([]string{"FreezeNodeResponse", string(data)}, " ")
}
