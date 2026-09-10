package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartModelServiceResponse Response Object
type StartModelServiceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StartModelServiceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartModelServiceResponse struct{}"
	}

	return strings.Join([]string{"StartModelServiceResponse", string(data)}, " ")
}
