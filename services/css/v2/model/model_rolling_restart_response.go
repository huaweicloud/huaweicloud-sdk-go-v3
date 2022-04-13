package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RollingRestartResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RollingRestartResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RollingRestartResponse struct{}"
	}

	return strings.Join([]string{"RollingRestartResponse", string(data)}, " ")
}
