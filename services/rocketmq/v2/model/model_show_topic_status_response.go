package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTopicStatusResponse Response Object
type ShowTopicStatusResponse struct {

	// 最大偏移量。
	MaxOffset *int32 `json:"max_offset,omitempty"`

	// 最小偏移量。
	MinOffset *int32 `json:"min_offset,omitempty"`

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
