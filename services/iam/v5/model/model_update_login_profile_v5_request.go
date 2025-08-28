package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLoginProfileV5Request Request Object
type UpdateLoginProfileV5Request struct {

	// IAM用户ID。
	UserId string `json:"user_id"`

	Body *UpdateLoginProfileReqBody `json:"body,omitempty"`
}

func (o UpdateLoginProfileV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLoginProfileV5Request struct{}"
	}

	return strings.Join([]string{"UpdateLoginProfileV5Request", string(data)}, " ")
}
