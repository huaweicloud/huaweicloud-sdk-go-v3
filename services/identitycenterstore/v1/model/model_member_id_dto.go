package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MemberIdDto An object containing the identifier of a group member.
type MemberIdDto struct {

	// 身份源中IdentityCenter用户的全局唯一标识符（ID）
	UserId string `json:"user_id"`
}

func (o MemberIdDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MemberIdDto struct{}"
	}

	return strings.Join([]string{"MemberIdDto", string(data)}, " ")
}
