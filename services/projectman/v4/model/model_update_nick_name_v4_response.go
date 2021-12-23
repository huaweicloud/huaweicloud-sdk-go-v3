package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateNickNameV4Response struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateNickNameV4Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNickNameV4Response struct{}"
	}

	return strings.Join([]string{"UpdateNickNameV4Response", string(data)}, " ")
}
