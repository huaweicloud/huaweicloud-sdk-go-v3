package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateServerNameRequest struct {

	// 服务器id。
	ServerId string `json:"server_id"`

	Body *UpdateServerNameRequestBody `json:"body,omitempty"`
}

func (o UpdateServerNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateServerNameRequest struct{}"
	}

	return strings.Join([]string{"UpdateServerNameRequest", string(data)}, " ")
}
