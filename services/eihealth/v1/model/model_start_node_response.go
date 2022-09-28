package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type StartNodeResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StartNodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartNodeResponse struct{}"
	}

	return strings.Join([]string{"StartNodeResponse", string(data)}, " ")
}
