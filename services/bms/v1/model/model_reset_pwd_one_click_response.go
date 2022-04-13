package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ResetPwdOneClickResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ResetPwdOneClickResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetPwdOneClickResponse struct{}"
	}

	return strings.Join([]string{"ResetPwdOneClickResponse", string(data)}, " ")
}
