package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetActivecodeResponse Response Object
type ResetActivecodeResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ResetActivecodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetActivecodeResponse struct{}"
	}

	return strings.Join([]string{"ResetActivecodeResponse", string(data)}, " ")
}
