package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateProcessInstanceReqOpinions struct {

	// 用户ID
	UserId *string `json:"user_id,omitempty"`

	// 当前责任人
	CurrOwner *string `json:"curr_owner,omitempty"`
}

func (o CreateProcessInstanceReqOpinions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProcessInstanceReqOpinions struct{}"
	}

	return strings.Join([]string{"CreateProcessInstanceReqOpinions", string(data)}, " ")
}
