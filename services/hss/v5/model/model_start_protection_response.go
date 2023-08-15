package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartProtectionResponse Response Object
type StartProtectionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StartProtectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartProtectionResponse struct{}"
	}

	return strings.Join([]string{"StartProtectionResponse", string(data)}, " ")
}
