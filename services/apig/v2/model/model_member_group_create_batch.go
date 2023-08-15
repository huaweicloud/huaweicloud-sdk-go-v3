package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MemberGroupCreateBatch struct {

	// 后端服务器组列表
	MemberGroups *[]MemberGroupCreate `json:"member_groups,omitempty"`
}

func (o MemberGroupCreateBatch) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MemberGroupCreateBatch struct{}"
	}

	return strings.Join([]string{"MemberGroupCreateBatch", string(data)}, " ")
}
