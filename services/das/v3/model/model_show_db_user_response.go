package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowDbUserResponse struct {
	DbUser         *DbUser `json:"db_user,omitempty" xml:"db_user"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDbUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDbUserResponse struct{}"
	}

	return strings.Join([]string{"ShowDbUserResponse", string(data)}, " ")
}
