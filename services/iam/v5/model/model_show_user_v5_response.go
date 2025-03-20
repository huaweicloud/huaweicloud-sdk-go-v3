package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowUserV5Response Response Object
type ShowUserV5Response struct {
	User           *UserEx `json:"user,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowUserV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUserV5Response struct{}"
	}

	return strings.Join([]string{"ShowUserV5Response", string(data)}, " ")
}
