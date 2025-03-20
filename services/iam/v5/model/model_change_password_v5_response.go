package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangePasswordV5Response Response Object
type ChangePasswordV5Response struct {
	HttpStatusCode int `json:"-"`
}

func (o ChangePasswordV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangePasswordV5Response struct{}"
	}

	return strings.Join([]string{"ChangePasswordV5Response", string(data)}, " ")
}
