package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartTaskResponse Response Object
type StartTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StartTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartTaskResponse struct{}"
	}

	return strings.Join([]string{"StartTaskResponse", string(data)}, " ")
}
