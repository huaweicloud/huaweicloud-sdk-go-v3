package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdatePasswordResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdatePasswordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePasswordResponse struct{}"
	}

	return strings.Join([]string{"UpdatePasswordResponse", string(data)}, " ")
}
