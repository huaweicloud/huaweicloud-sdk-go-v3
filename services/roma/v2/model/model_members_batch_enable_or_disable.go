package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MembersBatchEnableOrDisable struct {

	// 后端服务器编号列表。
	MemberIds *[]string `json:"member_ids,omitempty"`
}

func (o MembersBatchEnableOrDisable) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MembersBatchEnableOrDisable struct{}"
	}

	return strings.Join([]string{"MembersBatchEnableOrDisable", string(data)}, " ")
}
