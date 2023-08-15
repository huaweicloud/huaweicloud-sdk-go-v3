package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImportUserReq struct {

	// IAM用户id
	IamUserId string `json:"iam_user_id"`

	// 角色类型：管理员(ADMIN)、操作者(OPERATOR)
	Role string `json:"role"`

	Settings *UserSettingDto `json:"settings,omitempty"`
}

func (o ImportUserReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportUserReq struct{}"
	}

	return strings.Join([]string{"ImportUserReq", string(data)}, " ")
}
