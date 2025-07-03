package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetMmrLiveResponse Response Object
type SetMmrLiveResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SetMmrLiveResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetMmrLiveResponse struct{}"
	}

	return strings.Join([]string{"SetMmrLiveResponse", string(data)}, " ")
}
