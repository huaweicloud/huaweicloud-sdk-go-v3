package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartHouseKeepingResponse Response Object
type StartHouseKeepingResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StartHouseKeepingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartHouseKeepingResponse struct{}"
	}

	return strings.Join([]string{"StartHouseKeepingResponse", string(data)}, " ")
}
