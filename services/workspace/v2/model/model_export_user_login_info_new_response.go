package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportUserLoginInfoNewResponse Response Object
type ExportUserLoginInfoNewResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ExportUserLoginInfoNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportUserLoginInfoNewResponse struct{}"
	}

	return strings.Join([]string{"ExportUserLoginInfoNewResponse", string(data)}, " ")
}
