package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateDedicatedHostTagsRequest Request Object
type BatchCreateDedicatedHostTagsRequest struct {

	// 专属主机ID。  可以从专属主机控制台查询，或者通过调用查询专属主机列表API获取。
	DedicatedHostId string `json:"dedicated_host_id"`

	Body *ReqSetOrDeleteTags `json:"body,omitempty"`
}

func (o BatchCreateDedicatedHostTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateDedicatedHostTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateDedicatedHostTagsRequest", string(data)}, " ")
}
