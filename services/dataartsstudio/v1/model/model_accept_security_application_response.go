package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AcceptSecurityApplicationResponse Response Object
type AcceptSecurityApplicationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AcceptSecurityApplicationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AcceptSecurityApplicationResponse struct{}"
	}

	return strings.Join([]string{"AcceptSecurityApplicationResponse", string(data)}, " ")
}
