package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopInstanceResponse Response Object
type StopInstanceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StopInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopInstanceResponse struct{}"
	}

	return strings.Join([]string{"StopInstanceResponse", string(data)}, " ")
}
