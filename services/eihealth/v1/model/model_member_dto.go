package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MemberDto struct {

	// 用户id
	UserId string `json:"user_id"`
}

func (o MemberDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MemberDto struct{}"
	}

	return strings.Join([]string{"MemberDto", string(data)}, " ")
}
