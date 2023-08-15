package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDedicatedHostResponse Response Object
type ShowDedicatedHostResponse struct {
	DedicatedHost  *RespDedicatedHost `json:"dedicated_host,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ShowDedicatedHostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDedicatedHostResponse struct{}"
	}

	return strings.Join([]string{"ShowDedicatedHostResponse", string(data)}, " ")
}
