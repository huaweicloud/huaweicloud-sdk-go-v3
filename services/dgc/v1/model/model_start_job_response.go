package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type StartJobResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StartJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartJobResponse struct{}"
	}

	return strings.Join([]string{"StartJobResponse", string(data)}, " ")
}
