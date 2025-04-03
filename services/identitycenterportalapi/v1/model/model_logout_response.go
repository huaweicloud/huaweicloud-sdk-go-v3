package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LogoutResponse Response Object
type LogoutResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o LogoutResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogoutResponse struct{}"
	}

	return strings.Join([]string{"LogoutResponse", string(data)}, " ")
}
