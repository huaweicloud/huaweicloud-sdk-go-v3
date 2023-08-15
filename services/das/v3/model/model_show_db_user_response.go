package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDbUserResponse Response Object
type ShowDbUserResponse struct {
	DbUser         *DbUser `json:"db_user,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDbUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDbUserResponse struct{}"
	}

	return strings.Join([]string{"ShowDbUserResponse", string(data)}, " ")
}
