package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeIegPasswordRequest Request Object
type ChangeIegPasswordRequest struct {

	// 智能企业网关ID
	IegId string `json:"ieg_id"`

	Body *ChangePasswordBody `json:"body,omitempty"`
}

func (o ChangeIegPasswordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeIegPasswordRequest struct{}"
	}

	return strings.Join([]string{"ChangeIegPasswordRequest", string(data)}, " ")
}
