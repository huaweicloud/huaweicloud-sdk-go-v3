package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RemoveUsersResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RemoveUsersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveUsersResponse struct{}"
	}

	return strings.Join([]string{"RemoveUsersResponse", string(data)}, " ")
}
