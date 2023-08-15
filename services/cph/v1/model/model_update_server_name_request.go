package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateServerNameRequest Request Object
type UpdateServerNameRequest struct {

	// 云手机服务器的唯一标识。
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
