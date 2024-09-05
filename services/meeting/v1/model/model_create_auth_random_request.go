package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAuthRandomRequest Request Object
type CreateAuthRandomRequest struct {

	// 会议ID
	ConfId string `json:"conf_id"`

	// 0-不支持来宾会前等待页能力（默认）、1-支持来宾会前等待页能力
	GuestWaiting *int32 `json:"guest_waiting,omitempty"`

	// 会议密码
	XPassword *string `json:"X-Password,omitempty"`
}

func (o CreateAuthRandomRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAuthRandomRequest struct{}"
	}

	return strings.Join([]string{"CreateAuthRandomRequest", string(data)}, " ")
}
