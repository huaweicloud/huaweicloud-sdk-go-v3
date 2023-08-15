package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartSystemEventResponse Response Object
type StartSystemEventResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StartSystemEventResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartSystemEventResponse struct{}"
	}

	return strings.Join([]string{"StartSystemEventResponse", string(data)}, " ")
}
