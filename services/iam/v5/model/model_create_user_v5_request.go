package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateUserV5Request Request Object
type CreateUserV5Request struct {
	Body *CreateUserReqBody `json:"body,omitempty"`
}

func (o CreateUserV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateUserV5Request struct{}"
	}

	return strings.Join([]string{"CreateUserV5Request", string(data)}, " ")
}
