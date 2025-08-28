package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowUserV5Request Request Object
type ShowUserV5Request struct {

	// IAM用户ID。
	UserId string `json:"user_id"`
}

func (o ShowUserV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUserV5Request struct{}"
	}

	return strings.Join([]string{"ShowUserV5Request", string(data)}, " ")
}
