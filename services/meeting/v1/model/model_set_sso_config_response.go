package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetSsoConfigResponse Response Object
type SetSsoConfigResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SetSsoConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetSsoConfigResponse struct{}"
	}

	return strings.Join([]string{"SetSsoConfigResponse", string(data)}, " ")
}
