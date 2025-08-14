package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableUserResponse Response Object
type EnableUserResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o EnableUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableUserResponse struct{}"
	}

	return strings.Join([]string{"EnableUserResponse", string(data)}, " ")
}
