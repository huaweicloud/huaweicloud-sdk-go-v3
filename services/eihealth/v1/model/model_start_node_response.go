package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartNodeResponse Response Object
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
