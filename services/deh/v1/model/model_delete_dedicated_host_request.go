package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteDedicatedHostRequest struct {
	// 专属主机ID。

	DedicatedHostId string `json:"dedicated_host_id"`
}

func (o DeleteDedicatedHostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDedicatedHostRequest struct{}"
	}

	return strings.Join([]string{"DeleteDedicatedHostRequest", string(data)}, " ")
}
