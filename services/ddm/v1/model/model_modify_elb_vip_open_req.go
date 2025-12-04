package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ModifyElbVipOpenReq struct {

	// 组id。
	GroupId *string `json:"group_id,omitempty"`

	// 新ip。
	NewIp *string `json:"new_ip,omitempty"`
}

func (o ModifyElbVipOpenReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyElbVipOpenReq struct{}"
	}

	return strings.Join([]string{"ModifyElbVipOpenReq", string(data)}, " ")
}
