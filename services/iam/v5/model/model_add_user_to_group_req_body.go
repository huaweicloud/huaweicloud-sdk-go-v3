package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddUserToGroupReqBody struct {

	// IAM用户ID。
	UserId string `json:"user_id"`
}

func (o AddUserToGroupReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddUserToGroupReqBody struct{}"
	}

	return strings.Join([]string{"AddUserToGroupReqBody", string(data)}, " ")
}
