package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateServerRequest Request Object
type UpdateServerRequest struct {

	// 服务器唯一标识。
	ServerId string `json:"server_id"`

	Body *UpdateServerReq `json:"body,omitempty"`
}

func (o UpdateServerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateServerRequest struct{}"
	}

	return strings.Join([]string{"UpdateServerRequest", string(data)}, " ")
}
