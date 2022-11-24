package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchDeleteUsersResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteUsersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteUsersResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteUsersResponse", string(data)}, " ")
}
