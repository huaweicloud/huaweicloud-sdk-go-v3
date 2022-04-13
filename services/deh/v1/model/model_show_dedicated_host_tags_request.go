package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowDedicatedHostTagsRequest struct {
	// 专属主机ID。  可以从专属主机控制台查询，或者通过调用查询专属主机列表API获取。

	DedicatedHostId string `json:"dedicated_host_id"`
}

func (o ShowDedicatedHostTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDedicatedHostTagsRequest struct{}"
	}

	return strings.Join([]string{"ShowDedicatedHostTagsRequest", string(data)}, " ")
}
