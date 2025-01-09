package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ModifyPolicyTargetReq struct {

	// 添加应用。
	Add *[]Target `json:"add,omitempty"`

	// 删除应用。
	Delete *[]Target `json:"delete,omitempty"`
}

func (o ModifyPolicyTargetReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyPolicyTargetReq struct{}"
	}

	return strings.Join([]string{"ModifyPolicyTargetReq", string(data)}, " ")
}
