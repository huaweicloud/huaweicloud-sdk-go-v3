package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateMessageV2Req struct {
	Message *CreateMessageDoV2 `json:"message" xml:"message"`

	// 组id
	GroupId *string `json:"group_id,omitempty" xml:"group_id"`
}

func (o CreateMessageV2Req) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMessageV2Req struct{}"
	}

	return strings.Join([]string{"CreateMessageV2Req", string(data)}, " ")
}
