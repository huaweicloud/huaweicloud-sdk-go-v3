package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReinstallServerRequest Request Object
type ReinstallServerRequest struct {

	// 服务器唯一标识。
	ServerId string `json:"server_id"`

	Body *ReinstallServerReq `json:"body,omitempty"`
}

func (o ReinstallServerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReinstallServerRequest struct{}"
	}

	return strings.Join([]string{"ReinstallServerRequest", string(data)}, " ")
}
