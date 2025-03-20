package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowUserLastLoginV5Response Response Object
type ShowUserLastLoginV5Response struct {
	UserLastLogin  *UserLastLogin `json:"user_last_login,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowUserLastLoginV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUserLastLoginV5Response struct{}"
	}

	return strings.Join([]string{"ShowUserLastLoginV5Response", string(data)}, " ")
}
