package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateUserStatusResponse Response Object
type BatchUpdateUserStatusResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchUpdateUserStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateUserStatusResponse struct{}"
	}

	return strings.Join([]string{"BatchUpdateUserStatusResponse", string(data)}, " ")
}
