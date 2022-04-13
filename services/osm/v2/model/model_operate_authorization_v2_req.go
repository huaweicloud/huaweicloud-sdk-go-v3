package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OperateAuthorizationV2Req struct {
	// 拒绝理由

	RejectReason *string `json:"reject_reason,omitempty"`
	// 组id

	GroupId *string `json:"group_id,omitempty"`
}

func (o OperateAuthorizationV2Req) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperateAuthorizationV2Req struct{}"
	}

	return strings.Join([]string{"OperateAuthorizationV2Req", string(data)}, " ")
}
