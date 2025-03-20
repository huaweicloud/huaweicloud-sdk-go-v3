package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateUserV5Response Response Object
type UpdateUserV5Response struct {
	User           *User `json:"user,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o UpdateUserV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateUserV5Response struct{}"
	}

	return strings.Join([]string{"UpdateUserV5Response", string(data)}, " ")
}
