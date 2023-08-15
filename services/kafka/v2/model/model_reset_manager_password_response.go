package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetManagerPasswordResponse Response Object
type ResetManagerPasswordResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ResetManagerPasswordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetManagerPasswordResponse struct{}"
	}

	return strings.Join([]string{"ResetManagerPasswordResponse", string(data)}, " ")
}
