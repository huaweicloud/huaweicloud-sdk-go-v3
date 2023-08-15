package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopCbhInstanceResponse Response Object
type StopCbhInstanceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StopCbhInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopCbhInstanceResponse struct{}"
	}

	return strings.Join([]string{"StopCbhInstanceResponse", string(data)}, " ")
}
