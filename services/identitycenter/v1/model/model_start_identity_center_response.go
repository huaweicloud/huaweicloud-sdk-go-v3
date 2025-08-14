package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartIdentityCenterResponse Response Object
type StartIdentityCenterResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StartIdentityCenterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartIdentityCenterResponse struct{}"
	}

	return strings.Join([]string{"StartIdentityCenterResponse", string(data)}, " ")
}
