package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteUserV5Response Response Object
type DeleteUserV5Response struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteUserV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteUserV5Response struct{}"
	}

	return strings.Join([]string{"DeleteUserV5Response", string(data)}, " ")
}
