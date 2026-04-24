package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClientSecret Client secret for OAuth2 application.
type ClientSecret struct {
}

func (o ClientSecret) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClientSecret struct{}"
	}

	return strings.Join([]string{"ClientSecret", string(data)}, " ")
}
