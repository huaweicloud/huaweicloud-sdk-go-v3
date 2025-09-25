package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartManualSecurityCheckResponse Response Object
type StartManualSecurityCheckResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StartManualSecurityCheckResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartManualSecurityCheckResponse struct{}"
	}

	return strings.Join([]string{"StartManualSecurityCheckResponse", string(data)}, " ")
}
