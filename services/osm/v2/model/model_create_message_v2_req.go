package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateMessageV2Req struct {
	Message *CreateMessageDoV2 `json:"message"`

	// 华为云IAM组id，操作查询同组其他工单时，该id必传
	GroupId *string `json:"group_id,omitempty"`
}

func (o CreateMessageV2Req) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMessageV2Req struct{}"
	}

	return strings.Join([]string{"CreateMessageV2Req", string(data)}, " ")
}
