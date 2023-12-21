package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetInstancePasswordResponse Response Object
type ResetInstancePasswordResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ResetInstancePasswordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetInstancePasswordResponse struct{}"
	}

	return strings.Join([]string{"ResetInstancePasswordResponse", string(data)}, " ")
}
