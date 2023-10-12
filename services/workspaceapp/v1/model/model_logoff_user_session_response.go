package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LogoffUserSessionResponse Response Object
type LogoffUserSessionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o LogoffUserSessionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogoffUserSessionResponse struct{}"
	}

	return strings.Join([]string{"LogoffUserSessionResponse", string(data)}, " ")
}
