package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchDeleteDedicatedHostTagsRequest struct {

	// 专属主机ID。  可以从专属主机控制台查询，或者通过调用查询专属主机列表API获取。
	DedicatedHostId string `json:"dedicated_host_id" xml:"dedicated_host_id"`

	Body *ReqSetOrDeleteTags `json:"body,omitempty" xml:"body"`
}

func (o BatchDeleteDedicatedHostTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteDedicatedHostTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteDedicatedHostTagsRequest", string(data)}, " ")
}
