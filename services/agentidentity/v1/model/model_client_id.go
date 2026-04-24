package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClientId Client ID for OAuth2 application.
type ClientId struct {
}

func (o ClientId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClientId struct{}"
	}

	return strings.Join([]string{"ClientId", string(data)}, " ")
}
