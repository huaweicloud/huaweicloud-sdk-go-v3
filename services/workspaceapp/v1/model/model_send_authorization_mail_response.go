package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SendAuthorizationMailResponse Response Object
type SendAuthorizationMailResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SendAuthorizationMailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SendAuthorizationMailResponse struct{}"
	}

	return strings.Join([]string{"SendAuthorizationMailResponse", string(data)}, " ")
}
