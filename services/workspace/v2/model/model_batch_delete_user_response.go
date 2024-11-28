package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteUserResponse Response Object
type BatchDeleteUserResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteUserResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteUserResponse", string(data)}, " ")
}
