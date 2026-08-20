package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateProcessInstanceReqCcbs struct {

	// 用户ID
	UserId *string `json:"user_id,omitempty"`
}

func (o CreateProcessInstanceReqCcbs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProcessInstanceReqCcbs struct{}"
	}

	return strings.Join([]string{"CreateProcessInstanceReqCcbs", string(data)}, " ")
}
