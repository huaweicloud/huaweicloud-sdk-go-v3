package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyKafkaPublicIpAccessSwitchResponse Response Object
type ModifyKafkaPublicIpAccessSwitchResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ModifyKafkaPublicIpAccessSwitchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyKafkaPublicIpAccessSwitchResponse struct{}"
	}

	return strings.Join([]string{"ModifyKafkaPublicIpAccessSwitchResponse", string(data)}, " ")
}
