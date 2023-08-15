package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartCbhInstanceResponse Response Object
type StartCbhInstanceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StartCbhInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartCbhInstanceResponse struct{}"
	}

	return strings.Join([]string{"StartCbhInstanceResponse", string(data)}, " ")
}
