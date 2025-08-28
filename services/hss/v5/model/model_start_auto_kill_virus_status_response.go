package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartAutoKillVirusStatusResponse Response Object
type StartAutoKillVirusStatusResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StartAutoKillVirusStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartAutoKillVirusStatusResponse struct{}"
	}

	return strings.Join([]string{"StartAutoKillVirusStatusResponse", string(data)}, " ")
}
