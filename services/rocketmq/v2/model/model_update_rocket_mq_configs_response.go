package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRocketMqConfigsResponse Response Object
type UpdateRocketMqConfigsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateRocketMqConfigsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRocketMqConfigsResponse struct{}"
	}

	return strings.Join([]string{"UpdateRocketMqConfigsResponse", string(data)}, " ")
}
