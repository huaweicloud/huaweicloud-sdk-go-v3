package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowUserLastLoginV5Request Request Object
type ShowUserLastLoginV5Request struct {

	// IAM用户ID。
	UserId string `json:"user_id"`
}

func (o ShowUserLastLoginV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUserLastLoginV5Request struct{}"
	}

	return strings.Join([]string{"ShowUserLastLoginV5Request", string(data)}, " ")
}
