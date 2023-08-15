package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopAutoJobResponse Response Object
type StopAutoJobResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StopAutoJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopAutoJobResponse struct{}"
	}

	return strings.Join([]string{"StopAutoJobResponse", string(data)}, " ")
}
