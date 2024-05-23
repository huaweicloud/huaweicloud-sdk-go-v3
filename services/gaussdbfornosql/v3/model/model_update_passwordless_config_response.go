package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePasswordlessConfigResponse Response Object
type UpdatePasswordlessConfigResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdatePasswordlessConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePasswordlessConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdatePasswordlessConfigResponse", string(data)}, " ")
}
