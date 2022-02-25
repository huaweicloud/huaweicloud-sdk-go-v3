package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type MuteMeetingResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o MuteMeetingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MuteMeetingResponse struct{}"
	}

	return strings.Join([]string{"MuteMeetingResponse", string(data)}, " ")
}
