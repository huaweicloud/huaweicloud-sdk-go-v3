package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLoginProfileV5Request Request Object
type CreateLoginProfileV5Request struct {

	// IAM用户ID。
	UserId string `json:"user_id"`

	Body *CreateLoginProfileReqBody `json:"body,omitempty"`
}

func (o CreateLoginProfileV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLoginProfileV5Request struct{}"
	}

	return strings.Join([]string{"CreateLoginProfileV5Request", string(data)}, " ")
}
