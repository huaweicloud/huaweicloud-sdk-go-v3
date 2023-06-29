package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTopicStatusResponse Response Object
type ShowTopicStatusResponse struct {

	// 代理。
	Brokers        *[]ShowTopicStatusRespBrokers `json:"brokers,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ShowTopicStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTopicStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowTopicStatusResponse", string(data)}, " ")
}
