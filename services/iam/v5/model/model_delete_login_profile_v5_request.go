package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteLoginProfileV5Request Request Object
type DeleteLoginProfileV5Request struct {

	// IAM用户ID。
	UserId string `json:"user_id"`
}

func (o DeleteLoginProfileV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLoginProfileV5Request struct{}"
	}

	return strings.Join([]string{"DeleteLoginProfileV5Request", string(data)}, " ")
}
