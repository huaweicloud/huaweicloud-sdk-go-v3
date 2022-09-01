package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowDedicatedHostResponse struct {
	DedicatedHost  *RespDedicatedHost `json:"dedicated_host,omitempty" xml:"dedicated_host"`
	HttpStatusCode int                `json:"-"`
}

func (o ShowDedicatedHostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDedicatedHostResponse struct{}"
	}

	return strings.Join([]string{"ShowDedicatedHostResponse", string(data)}, " ")
}
