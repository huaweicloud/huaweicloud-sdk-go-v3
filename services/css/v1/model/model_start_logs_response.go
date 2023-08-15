package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartLogsResponse Response Object
type StartLogsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StartLogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartLogsResponse struct{}"
	}

	return strings.Join([]string{"StartLogsResponse", string(data)}, " ")
}
