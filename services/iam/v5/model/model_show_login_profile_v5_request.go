package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLoginProfileV5Request Request Object
type ShowLoginProfileV5Request struct {

	// IAM用户ID。
	UserId string `json:"user_id"`
}

func (o ShowLoginProfileV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLoginProfileV5Request struct{}"
	}

	return strings.Join([]string{"ShowLoginProfileV5Request", string(data)}, " ")
}
