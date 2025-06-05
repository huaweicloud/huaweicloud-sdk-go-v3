package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InitializePlatformResponse Response Object
type InitializePlatformResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o InitializePlatformResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InitializePlatformResponse struct{}"
	}

	return strings.Join([]string{"InitializePlatformResponse", string(data)}, " ")
}
