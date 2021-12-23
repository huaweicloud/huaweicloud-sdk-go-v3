package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateDedicatedHostRequest struct {
	// 专属主机ID。  可以从专属主机控制台查询，或者通过调用查询专属主机列表API获取。

	DedicatedHostId string `json:"dedicated_host_id"`

	Body *ReqUpdateDeh `json:"body,omitempty"`
}

func (o UpdateDedicatedHostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDedicatedHostRequest struct{}"
	}

	return strings.Join([]string{"UpdateDedicatedHostRequest", string(data)}, " ")
}
